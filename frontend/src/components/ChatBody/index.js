import React, { Component } from "react";

import { sendMsg, connect } from "./../../api";
import Message from "../Message";
import Intro from "../Intro";
import "./index.scss";

class ChatBody extends Component {
  constructor(props) {
    super(props)

    this.state = {
      chatHistory: [],
    };
  }

  componentDidMount() {
    connect((msg) => {
      this.setState(prevState => ({
        chatHistory: [...this.state.chatHistory, msg]
      }))
    })
  }

  handleSubmitMessage(event) {
    if (event.keyCode === 13) {
      sendMsg(event.target.value)
      event.target.value = ""
    } 
  }

  renderChatContent() {
    const { chatHistory } = this.state;
    return chatHistory.map(msg => <Message message={msg.data}/>);
  }

  renderChatInput() {
    return(
     <div>
       <input onKeyDown={this.handleSubmitMessage} />
     </div>
    )
  }

  renderChatBody() {
    return(
      <div className="chat-body">
        {this.renderChatContent()}
        {this.renderChatInput()}
      </div>
    )
  }

  render() {
    return(
      <div className="body">
        {/* <div className="welcome-body">
          <div> Welcome, member of Team</div>
        </div> */}
        <Intro/>
        {this.renderChatBody()}
      </div>
    )
  }
}

export default ChatBody