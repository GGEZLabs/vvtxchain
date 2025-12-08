package e2e

import (
	"fmt"
	"time"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (s *IntegrationTestSuite) testBankTokenTransfer() {
	s.Run("send_tokens_between_accounts", func() {
		var (
			err           error
			valIdx        = 0
			c             = s.chainA
			chainEndpoint = fmt.Sprintf("http://%s", s.valResources[c.id][valIdx].GetHostPort("1317/tcp"))
		)

		// define one sender and two recipient accounts
		alice, _ := c.genesisAccounts[1].keyInfo.GetAddress()
		bob, _ := c.genesisAccounts[2].keyInfo.GetAddress()
		charlie, _ := c.genesisAccounts[3].keyInfo.GetAddress()

		var beforeAliceUVVTXBalance,
			beforeBobUVVTXBalance,
			beforeCharlieUVVTXBalance,
			afterAliceUVVTXBalance,
			afterBobUVVTXBalance,
			afterCharlieUVVTXBalance sdk.Coin

		// get balances of sender and recipient accounts
		s.Require().Eventually(
			func() bool {
				beforeAliceUVVTXBalance, err = getSpecificBalance(chainEndpoint, alice.String(), uvvtxDenom)
				s.Require().NoError(err)

				beforeBobUVVTXBalance, err = getSpecificBalance(chainEndpoint, bob.String(), uvvtxDenom)
				s.Require().NoError(err)

				beforeCharlieUVVTXBalance, err = getSpecificBalance(chainEndpoint, charlie.String(), uvvtxDenom)
				s.Require().NoError(err)

				return beforeAliceUVVTXBalance.IsValid() && beforeBobUVVTXBalance.IsValid() && beforeCharlieUVVTXBalance.IsValid()
			},
			10*time.Second,
			5*time.Second,
		)

		// alice sends tokens to bob
		s.execBankSend(s.chainA, valIdx, alice.String(), bob.String(), tokenAmount.String(), standardFees.String(), false)

		// check that the transfer was successful
		s.Require().Eventually(
			func() bool {
				afterAliceUVVTXBalance, err = getSpecificBalance(chainEndpoint, alice.String(), uvvtxDenom)
				s.Require().NoError(err)

				afterBobUVVTXBalance, err = getSpecificBalance(chainEndpoint, bob.String(), uvvtxDenom)
				s.Require().NoError(err)

				// gasFeesBurnt := standardFees.Sub(sdk.NewCoin(UVVTXDenom, math.NewInt(1000)))
				// alice's balance should be decremented by the token amount and the gas fees
				// if the difference between expected and actual balance is less than 500, consider it as a success
				// any small change in operation/code can result in the gasFee difference
				// we set the threshold to 500 to avoid false negatives
				expectedAfterAliceUVVTXBalance := beforeAliceUVVTXBalance.Sub(tokenAmount).Sub(standardFees)
				decremented := afterAliceUVVTXBalance.Sub(expectedAfterAliceUVVTXBalance).Amount.LT(math.NewInt(500))

				incremented := beforeBobUVVTXBalance.Add(tokenAmount).IsEqual(afterBobUVVTXBalance)

				return decremented && incremented
			},
			10*time.Second,
			5*time.Second,
		)

		// save the updated account balances of alice and bob
		beforeAliceUVVTXBalance, beforeBobUVVTXBalance = afterAliceUVVTXBalance, afterBobUVVTXBalance

		// alice sends tokens to bob and charlie, at once
		s.execBankMultiSend(s.chainA, valIdx, alice.String(), []string{bob.String(), charlie.String()}, tokenAmount.String(), standardFees.String(), false)

		s.Require().Eventually(
			func() bool {
				afterAliceUVVTXBalance, err = getSpecificBalance(chainEndpoint, alice.String(), uvvtxDenom)
				s.Require().NoError(err)

				afterBobUVVTXBalance, err = getSpecificBalance(chainEndpoint, bob.String(), uvvtxDenom)
				s.Require().NoError(err)

				afterCharlieUVVTXBalance, err = getSpecificBalance(chainEndpoint, charlie.String(), uvvtxDenom)
				s.Require().NoError(err)

				// gasFeesBurnt := standardFees.Sub(sdk.NewCoin(UVVTXDenom, math.NewInt(1016)))
				// alice's balance should be decremented by the token amount and the gas fees
				// if the difference between expected and actual balance is less than 500, consider it as a success
				// any small change in operation/code can result in the gasFee difference
				// we set the threshold to 500 to avoid false negatives
				expectedAfterAliceUVVTXBalance := beforeAliceUVVTXBalance.Sub(tokenAmount).Sub(tokenAmount).Sub(standardFees)
				decremented := afterAliceUVVTXBalance.Sub(expectedAfterAliceUVVTXBalance).Amount.LT(math.NewInt(500))

				incremented := beforeBobUVVTXBalance.Add(tokenAmount).IsEqual(afterBobUVVTXBalance) &&
					beforeCharlieUVVTXBalance.Add(tokenAmount).IsEqual(afterCharlieUVVTXBalance)

				return decremented && incremented
			},
			10*time.Second,
			5*time.Second,
		)
	})
}
