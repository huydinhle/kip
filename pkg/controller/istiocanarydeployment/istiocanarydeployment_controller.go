package istiocanarydeployment

import (
	"context"

	"fmt"
	logr "github.com/go-logr/logr"
	appv1alpha1 "github.com/huydinhle/kip/pkg/apis/app/v1alpha1"
	istionetworking "istio.io/api/networking/v1alpha3"
	// istiov1alpha3 "istio.io/api/kube/apis/networking/v1alpha3"
	// istioclient "github.com/aspenmesh/istio-client-go/pkg/client/clientset/versioned"
	// istiov1alpha3 "istio.io/api/networking/v1alpha3"
	"istio.io/istio/pilot/pkg/config/kube/crd"
	"istio.io/istio/pilot/pkg/model"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	rest "k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_istiocanarydeployment")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new IstioCanaryDeployment Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileIstioCanaryDeployment{client: mgr.GetClient(), scheme: mgr.GetScheme(), config: mgr.GetConfig()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("istiocanarydeployment-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource IstioCanaryDeployment, make sure it should just handle CREATE and UPDATE
	err = c.Watch(&source.Kind{Type: &appv1alpha1.IstioCanaryDeployment{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner IstioCanaryDeployment

	// We are going to comment this part out because you don't really want to care about the deployment changes and stuff
	// err = c.Watch(&source.Kind{Type: &appsv1.Deployment{}}, &handler.EnqueueRequestForOwner{
	// 	IsController: true,
	// 	OwnerType:    &appv1alpha1.IstioCanaryDeployment{},
	// })
	// if err != nil {
	// 	return err
	// }

	// err = c.Watch(&source.Kind{Type: &corev1.Service{}}, &handler.EnqueueRequestForOwner{
	// 	IsController: true,
	// 	OwnerType:    &appv1alpha1.IstioCanaryDeployment{},
	// })
	// if err != nil {
	// 	return err
	// }

	return nil
}

var _ reconcile.Reconciler = &ReconcileIstioCanaryDeployment{}

// ReconcileIstioCanaryDeployment reconciles a IstioCanaryDeployment object
type ReconcileIstioCanaryDeployment struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
	config *rest.Config
}

// Reconcile reads that state of the cluster for a IstioCanaryDeployment object and makes changes based on the state read
// and what is in the IstioCanaryDeployment.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileIstioCanaryDeployment) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	fmt.Println("How many times this shit got called man")
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling IstioCanaryDeployment")

	// Fetch the IstioCanaryDeployment instance
	instance := &appv1alpha1.IstioCanaryDeployment{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	// Do regular check for our instance
	err = validateIstioCanaryDeploymentInstance(instance)
	if err != nil {
		return reconcile.Result{}, err
	}

	// define deployment, service, istio_vs materials
	deployment := newDeploymentForCR(instance, false)
	service := newServiceForCR(instance, false)
	canaryDeployment := newDeploymentForCR(instance, true)
	canaryService := newServiceForCR(instance, true)

	foundDeployment := &appsv1.Deployment{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: deployment.Name, Namespace: deployment.Namespace}, foundDeployment)

	if err != nil && errors.IsNotFound(err) {
		// if deployment DOES NOT exist flow
		reqLogger.Info("Creating a new Deployment, Service", "Deployment.Namespace", deployment.Namespace, "Deployment.Name", deployment.Name)

		//create Deployment and Service first
		err = createDeploymentAndService(r.client, instance, reqLogger, service, deployment, r.scheme)
		if err != nil {
			return reconcile.Result{}, err
		}

		// //update VirtualService
		// istioClient, err := istioclient.NewForConfig(r.config)
		// if err != nil {
		// 	return reconcile.Result{}, err
		// }

		err = updateVirtualService(instance, reqLogger)
		if err != nil {
			return reconcile.Result{}, err
		}

		return reconcile.Result{}, nil

	} else if err != nil {
		return reconcile.Result{}, err
	}

	// if deployment DOES exist, create Canary Materials
	err = createCanaryMaterials(r.client, instance, reqLogger, canaryDeployment, canaryService, r.scheme)
	if err != nil {
		return reconcile.Result{}, err
	}

	return reconcile.Result{}, nil
}

// We need to work on this, we need to steal this somewhere to do validation for our Canary before
// installing it in our cluster
func validateIstioCanaryDeploymentInstance(instance *appv1alpha1.IstioCanaryDeployment) error {
	return nil
}

func createCanaryMaterials(client client.Client, instance *appv1alpha1.IstioCanaryDeployment, reqLogger logr.Logger, canaryDeployment *appsv1.Deployment, canaryService *corev1.Service, scheme *runtime.Scheme) error {
	// Delete existing Canary Deployment
	err := deleteDeployment(client, reqLogger, canaryDeployment)
	if err != nil {
		return err
	}

	// Delete existing Canary Service
	err = deleteService(client, reqLogger, canaryService)
	if err != nil {
		return err
	}

	//create Deployment and Service for Canary
	err = createDeploymentAndService(client, instance, reqLogger, canaryService, canaryDeployment, scheme)
	if err != nil {
		return err
	}

	return nil
}

// newDeploymentForCR returns a busybox pod with the same name/namespace as the cr
func newDeploymentForCR(cr *appv1alpha1.IstioCanaryDeployment, canary bool) *appsv1.Deployment {
	name := cr.Name
	if canary {
		name = name + "-canary"
	}

	labels := map[string]string{
		"app": name,
	}

	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name + "-deployment",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: cr.Spec.DeploymentSpec,
	}
}

// newDeploymentForCR returns a busybox pod with the same name/namespace as the cr
func newServiceForCR(cr *appv1alpha1.IstioCanaryDeployment, canary bool) *corev1.Service {
	name := cr.Name
	if canary {
		name = name + "-canary"
	}

	labels := map[string]string{
		"app": name,
	}

	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name + "-service",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: cr.Spec.ServiceSpec,
	}
}

func createDeploymentAndService(client client.Client, instance *appv1alpha1.IstioCanaryDeployment, reqLogger logr.Logger, service *corev1.Service, deployment *appsv1.Deployment, scheme *runtime.Scheme) error {
	reqLogger.Info("Creating a new Deployment and Service", "Namespace", deployment.Namespace, "Name", deployment.Name)

	if err := controllerutil.SetControllerReference(instance, deployment, scheme); err != nil {
		return err
	}

	if err := controllerutil.SetControllerReference(instance, service, scheme); err != nil {
		return err
	}

	if err := client.Create(context.TODO(), service); err != nil {
		return err
	}

	if err := client.Create(context.TODO(), deployment); err != nil {
		return err
	}

	return nil
}

func updateVirtualService(instance *appv1alpha1.IstioCanaryDeployment, reqLogger logr.Logger) error {

	reqLogger.Info("Updating Service Flow")
	client, err := crd.NewClient("", "", model.IstioConfigTypes, "")
	if err != nil {
		return err
	}

	// client.Get(typ, name, namespace string) (config *Config, exists bool)
	vsConfig, exist := client.Get(model.VirtualService.Type, instance.Spec.VSName, instance.Namespace)
	if !exist {
		return fmt.Errorf("the virtual service %s didn't exist", instance.Spec.VSName)
	}

	vs, ok := vsConfig.Spec.(*istionetworking.VirtualService)
	if !ok { // should never happen
		return fmt.Errorf("in not a virtual service: %#v", vsConfig)
	}

	// testing for fun
	vs.Http[0].Timeout.Seconds = 120

	if vsConfig.ConfigMeta.Labels == nil {
		vsConfig.ConfigMeta.Labels = make(map[string]string)
	}
	vsConfig.ConfigMeta.Labels["kip"] = "canary"
	_, err = client.Update(*vsConfig)
	if err != nil {
		return fmt.Errorf("Can't update the virtual service %s", vsConfig.ConfigMeta.Name)
	}
	return nil
}

func deleteService(client client.Client, reqLogger logr.Logger, service *corev1.Service) error {
	reqLogger.Info("Deleting Service", "Service.Namespace", service.Namespace, "Service.Name", service.Name)

	foundService := &corev1.Service{}
	err := client.Get(context.TODO(), types.NamespacedName{Name: service.Name, Namespace: service.Namespace}, foundService)

	if err != nil && errors.IsNotFound(err) {
		return nil
	}

	if err := client.Delete(context.TODO(), service); err != nil {
		return err
	}

	return nil
}

func deleteDeployment(client client.Client, reqLogger logr.Logger, deployment *appsv1.Deployment) error {
	reqLogger.Info("Deleting Deployment", "Namespace", deployment.Namespace, "Name", deployment.Name)

	foundDeployment := &appsv1.Deployment{}
	err := client.Get(context.TODO(), types.NamespacedName{Name: deployment.Name, Namespace: deployment.Namespace}, foundDeployment)

	if err != nil && errors.IsNotFound(err) {
		return nil
	}

	if err := client.Delete(context.TODO(), deployment); err != nil {
		return err
	}

	return nil
}

// func deleteVirtualService(client client.Client, reqLogger logr.Logger, vs *istiov1alpha3.VirtualService) error {
// 	reqLogger.Info("Deleting VirtualService", "Namespace", vs.Namespace, "Name", vs.Name)

// 	foundVS := &istiov1alpha3.VirtualService{}
// 	err := client.Get(context.TODO(), types.NamespacedName{Name: vs.Name, Namespace: vs.Namespace}, foundVS)

// 	if err != nil && errors.IsNotFound(err) {
// 		return nil
// 	}

// 	if err := client.Delete(context.TODO(), vs); err != nil {
// 		return err
// 	}

// 	return nil
// }
