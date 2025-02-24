package cmd

import (
	"errors"
	"strings"
	"testing"

	"github.com/bmizerany/assert"
	"github.com/keptn/keptn/cli/pkg/common"
	commonfake "github.com/keptn/keptn/cli/pkg/common/fake"
	helmfake "github.com/keptn/keptn/cli/pkg/helm/fake"
	kubefake "github.com/keptn/keptn/cli/pkg/kube/fake"
	"helm.sh/helm/v3/pkg/chart"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestInstallCmdHandler_doInstallation(t *testing.T) {
	type fields struct {
		helmHelper       *helmfake.IHelperMock
		namespaceHandler *kubefake.IKeptnNamespaceHandlerMock
		userInput        *commonfake.IUserInputMock
	}
	tests := []struct {
		name              string
		fields            fields
		args              installCmdParams
		chartsToBeApplied []*chart.Chart
		wantErr           bool
	}{
		{
			name: "installation: continuous delivery",
			fields: fields{
				helmHelper: &helmfake.IHelperMock{
					DownloadChartFunc: func(chartRepoURL string) (*chart.Chart, error) {
						var chartName string
						if strings.Contains(chartRepoURL, helmServiceName) {
							chartName = helmServiceName
						} else if strings.Contains(chartRepoURL, jmeterServiceName) {
							chartName = jmeterServiceName
						} else {
							chartName = keptnReleaseName
						}
						return &chart.Chart{
							Metadata: &chart.Metadata{Name: chartName},
						}, nil
					},
					UpgradeChartFunc: func(ch *chart.Chart, releaseName string, namespace string, vals map[string]interface{}) error {
						return nil
					},
				},
				namespaceHandler: &kubefake.IKeptnNamespaceHandlerMock{
					CreateNamespaceFunc: func(useInClusterConfig bool, namespace string, namespaceMetadata ...metav1.ObjectMeta) error {
						return nil
					},
					ExistsNamespaceFunc: func(useInClusterConfig bool, namespace string) (bool, error) {
						return false, nil
					},
				},
				userInput: &commonfake.IUserInputMock{AskBoolFunc: func(question string, opts *common.UserInputOptions) bool {
					return true
				}},
			},
			args: installCmdParams{
				UseCase: ContinuousDelivery,
				installUpgradeParams: installUpgradeParams{
					PlatformIdentifier: stringp("kubernetes"),
				},
				UseCaseInput:      stringp("continuous-delivery"),
				HideSensitiveData: boolp(false),
			},
			chartsToBeApplied: []*chart.Chart{
				{
					Metadata: &chart.Metadata{Name: keptnReleaseName},
				},
				{
					Metadata: &chart.Metadata{Name: helmServiceName},
				},
				{
					Metadata: &chart.Metadata{Name: jmeterServiceName},
				},
			},
			wantErr: false,
		},
		{
			name: "installation: control plane only",
			fields: fields{
				helmHelper: &helmfake.IHelperMock{
					DownloadChartFunc: func(chartRepoURL string) (*chart.Chart, error) {
						var chartName string
						if strings.Contains(chartRepoURL, helmServiceName) {
							return nil, errors.New("should not be called")
						} else if strings.Contains(chartRepoURL, jmeterServiceName) {
							return nil, errors.New("should not be called")
						} else {
							chartName = keptnReleaseName
						}
						return &chart.Chart{
							Metadata: &chart.Metadata{Name: chartName},
						}, nil
					},
					UpgradeChartFunc: func(ch *chart.Chart, releaseName string, namespace string, vals map[string]interface{}) error {
						return nil
					},
				},
				namespaceHandler: &kubefake.IKeptnNamespaceHandlerMock{
					CreateNamespaceFunc: func(useInClusterConfig bool, namespace string, namespaceMetadata ...metav1.ObjectMeta) error {
						return nil
					},
					ExistsNamespaceFunc: func(useInClusterConfig bool, namespace string) (bool, error) {
						return false, nil
					},
				},
				userInput: &commonfake.IUserInputMock{AskBoolFunc: func(question string, opts *common.UserInputOptions) bool {
					return true
				}},
			},
			args: installCmdParams{
				UseCase: QualityGates,
				installUpgradeParams: installUpgradeParams{
					PlatformIdentifier: stringp("kubernetes"),
				},
				UseCaseInput:      stringp(""),
				HideSensitiveData: boolp(false),
			},
			chartsToBeApplied: []*chart.Chart{
				{
					Metadata: &chart.Metadata{Name: keptnReleaseName},
				},
			},
			wantErr: false,
		},
		{
			name: "installation: control plane only, custom chart URL",
			fields: fields{
				helmHelper: &helmfake.IHelperMock{
					DownloadChartFunc: func(chartRepoURL string) (*chart.Chart, error) {
						var chartName string
						if strings.Contains(chartRepoURL, helmServiceName) {
							return nil, errors.New("should not be called")
						} else if strings.Contains(chartRepoURL, jmeterServiceName) {
							return nil, errors.New("should not be called")
						} else {
							chartName = keptnReleaseName
						}
						return &chart.Chart{
							Metadata: &chart.Metadata{Name: chartName},
						}, nil
					},
					UpgradeChartFunc: func(ch *chart.Chart, releaseName string, namespace string, vals map[string]interface{}) error {
						return nil
					},
				},
				namespaceHandler: &kubefake.IKeptnNamespaceHandlerMock{
					CreateNamespaceFunc: func(useInClusterConfig bool, namespace string, namespaceMetadata ...metav1.ObjectMeta) error {
						return nil
					},
					ExistsNamespaceFunc: func(useInClusterConfig bool, namespace string) (bool, error) {
						return false, nil
					},
				},
				userInput: &commonfake.IUserInputMock{AskBoolFunc: func(question string, opts *common.UserInputOptions) bool {
					return true
				}},
			},
			args: installCmdParams{
				UseCase: QualityGates,
				installUpgradeParams: installUpgradeParams{
					PlatformIdentifier: stringp("kubernetes"),
					ChartRepoURL:       stringp("https://charts-dev.keptn.sh/packages/keptn-0.11.4.tgz"),
				},
				UseCaseInput:      stringp(""),
				HideSensitiveData: boolp(false),
			},
			chartsToBeApplied: []*chart.Chart{
				{
					Metadata: &chart.Metadata{Name: keptnReleaseName},
				},
			},
			wantErr: false,
		},
		{
			name: "installation: control plane only, namespace exists, cancel installation",
			fields: fields{
				helmHelper: &helmfake.IHelperMock{
					DownloadChartFunc: func(chartRepoURL string) (*chart.Chart, error) {
						var chartName string
						if strings.Contains(chartRepoURL, helmServiceName) {
							return nil, errors.New("should not be called")
						} else if strings.Contains(chartRepoURL, jmeterServiceName) {
							return nil, errors.New("should not be called")
						} else {
							chartName = keptnReleaseName
						}
						return &chart.Chart{
							Metadata: &chart.Metadata{Name: chartName},
						}, nil
					},
					UpgradeChartFunc: func(ch *chart.Chart, releaseName string, namespace string, vals map[string]interface{}) error {
						return nil
					},
				},
				namespaceHandler: &kubefake.IKeptnNamespaceHandlerMock{
					CreateNamespaceFunc: func(useInClusterConfig bool, namespace string, namespaceMetadata ...metav1.ObjectMeta) error {
						return nil
					},
					ExistsNamespaceFunc: func(useInClusterConfig bool, namespace string) (bool, error) {
						return true, nil
					},
				},
				userInput: &commonfake.IUserInputMock{AskBoolFunc: func(question string, opts *common.UserInputOptions) bool {
					return false
				}},
			},
			args: installCmdParams{
				UseCase: QualityGates,
				installUpgradeParams: installUpgradeParams{
					PlatformIdentifier: stringp("kubernetes"),
				},
				UseCaseInput:      stringp(""),
				HideSensitiveData: boolp(false),
			},
			chartsToBeApplied: []*chart.Chart{},
			wantErr:           true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &InstallCmdHandler{
				helmHelper:       tt.fields.helmHelper,
				namespaceHandler: tt.fields.namespaceHandler,
				userInput:        tt.fields.userInput,
			}
			if err := i.doInstallation(tt.args); (err != nil) != tt.wantErr {
				t.Errorf("doInstallation() error = %v, wantErr %v", err, tt.wantErr)
			}

			// has namespace been checked?
			assert.Equal(t, 1, len(tt.fields.namespaceHandler.ExistsNamespaceCalls()))

			for index, upgradeChartCall := range tt.fields.helmHelper.UpgradeChartCalls() {
				assert.Equal(t, tt.chartsToBeApplied[index], upgradeChartCall.Ch)
			}
		})
	}
}

// TestInstallUnknownCommand
func TestInstallUnknownCommand(t *testing.T) {
	testInvalidInputHelper("install someUnknownCommand", "unknown command \"someUnknownCommand\" for \"keptn install\"", t)
}

// TestInstallUnknownParameter
func TestInstallUnknownParmeter(t *testing.T) {
	testInvalidInputHelper("install --project=sockshop", "unknown flag: --project", t)
}
