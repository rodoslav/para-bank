// src/utils/cryptoUtils.js
import { generateMnemonic as generate } from 'bip39';

export const generateMnemonic = () => {
  return generate(128); // Генерація сім'ї слів з 128 біт ентропії
};
