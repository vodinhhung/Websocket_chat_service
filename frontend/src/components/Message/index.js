import React, { Component } from "react";

import "./index.scss";

class Message extends Component {
  constructor(props) {
    super(props);
  }

  deserialize(msg) {
    let temp = JSON.parse(msg)
    return temp.body
  }

  render() {
    const { message } = this.props;
    let msg = this.deserialize(message)

    return <div className="message"> {msg} </div> 
  }
}

export default Message