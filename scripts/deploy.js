const CashbackRewardsToken = artifacts.require("CashbackRewardsToken");
const RewardPointsToken = artifacts.require("RewardPointsToken");

module.exports = async function (deployer, network, accounts) {
  await deployer.deploy(CashbackRewardsToken);
  const cashbackRewardsToken = await CashbackRewardsToken.deployed();

  await deployer.deploy(RewardPointsToken);
  const rewardPointsToken = await RewardPointsToken.deployed();

  console.log("CashbackRewardsToken deployed at:", cashbackRewardsToken.address);
  console.log("RewardPointsToken deployed at:", rewardPointsToken.address);
};