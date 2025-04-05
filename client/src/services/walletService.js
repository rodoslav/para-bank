// src/services/walletService.js
import { generateMnemonic } from '../utils/cryptoUtils';

export const createWallet = () => {
  const seedPhrase = generateMnemonic();
  // Логіка для створення гаманця з seedPhrase
  return { seedPhrase };
};
