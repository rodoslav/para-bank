import React from 'react';
import CreateWallet from './components/CreateWallet';
import WalletBalance from './components/WalletBalance';
import SendTransaction from './components/SendTransaction';
import TransactionHistory from './components/TransactionHistory';
import Header from './components/Header';

const App = () => {
  return (
    <div className="App">
      <Header />
      <CreateWallet />
      <WalletBalance />
      <SendTransaction />
      <TransactionHistory />
    </div>
  );
};

export default App;
