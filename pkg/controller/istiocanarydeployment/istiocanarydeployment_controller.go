package istiocanarydeployment

import (
	"context"

	"fmt"
	logr "github.com/go-logr/logr"
	appv1alpha1 "github.com/huydinhle/kip/pkg/apis/app/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
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
	return &ReconcileIstioCanaryDeployment{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("istiocanarydeployment-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource IstioCanaryDeployment
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
}

// Reconcile reads that state of the cluster for a IstioCanaryDeployment object and makes changes based on the state read
// and what is in the IstioCanaryDeployment.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileIstioCanaryDeployment) Reconcile(request reconcile.Request) (reconcile.Result, error) {
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

	// Define a new Deployment object
	deployment := newDeploymentForCR(instance)
	service := newServiceForCR(instance)
	// fmt.Printf("deployment = %+v\n", deployment)
	fmt.Println("Debug")

	// Set CanaryDeployment instance as the owner and controller for Deployment and Service
	if err := controllerutil.SetControllerReference(instance, deployment, r.scheme); err != nil {
		return reconcile.Result{}, err
	}
	if err := controllerutil.SetControllerReference(instance, service, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	// handleNewService
	err = handleNewService(service, r.client, reqLogger)
	if err != nil {
		return reconcile.Result{}, err
	}

	// Handling Deployment
	err = handleNewDeployment(deployment, r.client, reqLogger, instance)
	if err != nil {
		return reconcile.Result{}, err
	}

	fmt.Println("\n")
	return reconcile.Result{}, nil
}

// We need to work on this, we need to steal this somewhere to do validation for our Canary before
// installing it in our cluster
func validateIstioCanaryDeploymentInstance(instance *appv1alpha1.IstioCanaryDeployment) error {
	return nil
}

func handleNewDeployment(deployment *appsv1.Deployment, client client.Client, reqLogger logr.Logger, cr *appv1alpha1.IstioCanaryDeployment) error {
	var err error
	foundDeployment := &appsv1.Deployment{}
	err = client.Get(context.TODO(), types.NamespacedName{Name: deployment.Name, Namespace: deployment.Namespace}, foundDeployment)

	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new Deployment", "Deployment.Namespace", deployment.Namespace, "Deployment.Name", deployment.Name)
		err = client.Create(context.TODO(), deployment)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	// Deployment already existed, put deployment into canary pipeline
	return canaryDeploymentHandling(client, cr, reqLogger)
}

func canaryDeploymentHandling(client client.Client, cr *appv1alpha1.IstioCanaryDeployment, reqLogger logr.Logger) error {
	var err error
	deployment := newCanaryDeploymentForCR(cr)
	//find if there is a canary running, if there is delete it
	foundDeployment := &appsv1.Deployment{}
	err = client.Get(context.TODO(), types.NamespacedName{Name: deployment.Name, Namespace: deployment.Namespace}, foundDeployment)

	// if canary foundDeployment exist delete it
	reqLogger.Info("Deleteing canary deployment ", "Deployment.Namespace", deployment.Namespace, "deployment.Name", deployment.Name)
	if foundDeployment.Name != "" {
		err = client.Delete(context.TODO(), foundDeployment)
		if err != nil {
			return err
		}
	}

	reqLogger.Info("Creating a new Deployment", "Deployment.Namespace", deployment.Namespace, "Deployment.Name", deployment.Name)
	err = client.Create(context.TODO(), deployment)
	if err != nil {
		return err
	}
	return nil
}

func handleNewService(service *corev1.Service, client client.Client, reqLogger logr.Logger) error {
	var err error
	foundService := &corev1.Service{}
	err = client.Get(context.TODO(), types.NamespacedName{Name: service.Name, Namespace: service.Namespace}, foundService)

	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new Service", "Service.Namespace", service.Namespace, "Service.Name", service.Name)
		err = client.Create(context.TODO(), service)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	// Service already existed so we don't do anything for now
	reqLogger.Info("Updating Service Skipped: Service already exists", "Service.Namespace", service.Namespace, "Service.Name", service.Name)

	return nil
}

// newDeploymentForCR returns a busybox pod with the same name/namespace as the cr
func newDeploymentForCR(cr *appv1alpha1.IstioCanaryDeployment) *appsv1.Deployment {
	labels := map[string]string{
		"app": cr.Name,
	}
	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-deployment",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: cr.Spec.DeploymentSpec,
	}
}

// newDeploymentForCR returns a busybox pod with the same name/namespace as the cr
func newServiceForCR(cr *appv1alpha1.IstioCanaryDeployment) *corev1.Service {
	labels := map[string]string{
		"app": cr.Name,
	}
	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-deployment",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: cr.Spec.ServiceSpec,
	}
}

// newDeploymentForCR returns a busybox pod with the same name/namespace as the cr
func newCanaryDeploymentForCR(cr *appv1alpha1.IstioCanaryDeployment) *appsv1.Deployment {
	labels := map[string]string{
		"app": cr.Name + "-canary",
	}
	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-deployment" + "-canary",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: cr.Spec.DeploymentSpec,
	}
}

// newDeploymentForCR returns a busybox pod with the same name/namespace as the cr
func newCanaryServiceForCR(cr *appv1alpha1.IstioCanaryDeployment) *corev1.Service {
	labels := map[string]string{
		"app": cr.Name + "-canary",
	}
	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-deployment" + "-canary",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: cr.Spec.ServiceSpec,
	}
}
