import { ethers } from "ethers";
import CashbackRewardsToken from "../build/CashbackRewardsToken.json";

export const getContract = async (provider) => {
  const signer = provider.getSigner();
  const contract = new ethers.Contract(
    "YOUR_CONTRACT_ADDRESS",
    CashbackRewardsToken.abi,
    signer
  );
  return contract;
};