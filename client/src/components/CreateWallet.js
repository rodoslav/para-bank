// src/components/CreateWallet.js
import React, { useState } from 'react';
import { createWallet } from '../services/walletService';

const CreateWallet = () => {
  const [seedPhrase, setSeedPhrase] = useState('');

  const handleCreateWallet = async () => {
    const wallet = await createWallet();
    setSeedPhrase(wallet.seedPhrase);
  };

  return (
    <div>
      <h2>Create New Wallet</h2>
      <button onClick={handleCreateWallet}>Create Wallet</button>
      {seedPhrase && (
        <div>
          <h3>Your Seed Phrase:</h3>
          <p>{seedPhrase}</p>
        </div>
      )}
    </div>
  );
};

export default CreateWallet;
