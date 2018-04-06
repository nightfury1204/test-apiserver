package install

import (
	"k8s.io/apimachinery/pkg/apimachinery/announced"
	"k8s.io/apimachinery/pkg/apimachinery/registered"
	"k8s.io/apimachinery/pkg/runtime"
	"github.com/nightfury1204/test-apiserver/pkg/apis/tryapi"
	"github.com/nightfury1204/test-apiserver/pkg/apis/tryapi/v1alpha1"
)

// Install registers the API group and adds types to a scheme
func Install(groupFactoryRegistry announced.APIGroupFactoryRegistry, registry *registered.APIRegistrationManager, scheme *runtime.Scheme) {
	if err := announced.NewGroupMetaFactory(
		&announced.GroupMetaFactoryArgs{
			GroupName:                  tryapi.GroupName,
			// RootScopedKinds:            sets.NewString("Fischer", "FischerList"),
			VersionPreferenceOrder:     []string{v1alpha1.SchemeGroupVersion.Version},
			AddInternalObjectsToScheme: tryapi.AddToScheme,
		},
		announced.VersionToSchemeFunc{
			v1alpha1.SchemeGroupVersion.Version: v1alpha1.AddToScheme,
		},
	).Announce(groupFactoryRegistry).RegisterAndEnable(registry, scheme); err != nil {
		panic(err)
	}
}
