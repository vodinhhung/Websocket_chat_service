import React, { Component } from 'react';

import Header from './components/Header';
import Footer from './components/Footer';
import ChatBody from './components/ChatBody';

class App extends Component {
  render() {
    return (
      <div className="App">
        <Header/>
        <ChatBody/>
        <Footer/>
      </div>
    );
  }
}

export default App;
