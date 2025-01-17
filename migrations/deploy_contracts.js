const CashbackRewardsToken = artifacts.require("CashbackRewardsToken");
const RewardPointsToken = artifacts.require("RewardPointsToken");

module.exports = async function (deployer, network, accounts) {
  await deployer.deploy(CashbackRewardsToken);
  const cashbackRewardsToken = await CashbackRewardsToken.deployed();

  console.log(`CashbackRewardsToken deployed at: ${cashbackRewardsToken.address}`);

  await deployer.deploy(RewardPointsToken);
  const rewardPointsToken = await RewardPointsToken.deployed();

  console.log(`RewardPointsToken deployed at: ${rewardPointsToken.address}`);
};