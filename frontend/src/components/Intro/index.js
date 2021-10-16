import React, { Component } from "react";

import "./index.scss";

class Intro extends Component {
  render() {
    return (
      <div className="welcome-body">
        <div className="avatar">
          <img src="./download.png"/>
        </div>
      </div>
    )
  }
}

export default Intro;