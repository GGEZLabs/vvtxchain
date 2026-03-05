package app

import (
	"fmt"

	"github.com/GGEZLabs/vvtxchain/app/upgrade/v1_0_1"
	"github.com/cosmos/cosmos-sdk/types/module"
)

func (app *App) setupUpgradeHandlers(configurator module.Configurator) {
	app.UpgradeKeeper.SetUpgradeHandler(
		v1_0_1.UpgradeName,
		v1_0_1.CreateUpgradeHandler(app.ModuleManager, configurator),
	)

	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(fmt.Errorf("failed to read upgrade info from disk: %w", err))
	}

	if app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		return
	}
}
