import React, { Component, useRef, useEffect, useState } from "react";
import "./index.css";
import useWebSocket, { ReadyState } from "react-use-websocket";
import { json } from "react-router-dom";
import { Navigate } from "react-router-dom";
import Cookies from "js-cookie";
import { jwtDecode } from "jwt-decode";

export default class Chat extends Component {
  constructor(props) {
    super();
    super.constructor(props);

    this.state = {
      authenticated: !(
        Cookies.get("token") == null || Cookies.get("token") == ""
      ),
      userId: null,
    };

    this.state.userId = `chatUser${Math.floor(Math.random() * 100 + 1)}`;

    console.log(this.state);

    if (this.state.authenticated) {
      let {role, user_id, exp} = jwtDecode(Cookies.get("token"));
      console.log("claims:", role, user_id, exp);
      this.state.userId = user_id;
    }
  }

  render() {
    if (!this.state.authenticated) {
      return <Navigate to="/sign-in" />;
    }

    return (
      <div className="card chat-app">
        <div id="plist" className="people-list">
          <ul className="list-unstyled chat-list mt-2 mb-0">
            <li className="clearfix">
              <img
                src="https://bootdey.com/img/Content/avatar/avatar1.png"
                alt="avatar"
              />
              <div className="about">
                <div className="name">Vincent Porter</div>
                <div className="status">
                  {" "}
                  <i className="fa fa-circle offline"></i> left 7 mins ago{" "}
                </div>
              </div>
            </li>
            <li className="clearfix active">
              <img
                src="https://bootdey.com/img/Content/avatar/avatar2.png"
                alt="avatar"
              />
              <div className="about">
                <div className="name">Aiden Chavez</div>
                <div className="status">
                  {" "}
                  <i className="fa fa-circle online"></i> online{" "}
                </div>
              </div>
            </li>
          </ul>
        </div>
        <div className="chat">
          <TextBx userId={this.state.userId} />
        </div>
      </div>
    );
  }
}

const useWs = (userId) => {
  const [isReady, setIsReady] = useState(false);
  const [val, setVal] = useState(null);

  const ws = useRef(null);

  const socket = new WebSocket(`ws://localhost:8080/v4/chat/${userId}`);

  useEffect(() => {
    socket.onopen = () => setIsReady(true);
    socket.onclose = () => setIsReady(false);
    socket.onmessage = (event) => setVal(event.data);

    ws.current = socket;

    return () => {
      // socket.close();
    };
  }, []);

  // bind is needed to make sure `send` references correct `this`
  return [isReady, val, ws.current?.send.bind(ws.current)];
};

const TextBx = ({ userId }) => {
  const [ready, val, send] = useWs(userId);

  // useEffect(() => {
  //   if (ready) {
  //     // send("test message")
  //   }
  // }, [ready, send]); // make sure to include send in dependency array

  let msgs = [];
  let dataFound = val != null;
  if (dataFound) {
    msgs = JSON.parse(val);
  }

  var handleSubmit = (event) => {
    event.preventDefault();
    send(event.target[0].value);
  };

  return (
    <form onSubmit={handleSubmit}>
      <div className="chat-history">
        <ul className="m-b-0">
          {msgs.map((msg) => {
            if (userId !== msg.UserId) {
              return (
                <li className="clearfix">
                  <div className="message my-message">
                    [{msg.UserId}] {msg.Text}
                    <br />
                    <span className="message-data-time">
                      {new Date(msg.Timestamp).toLocaleString()}
                    </span>
                  </div>
                </li>
              );
            } else {
              return (
                <li className="clearfix">
                  <div className="message other-message float-right">
                    [{msg.UserId}] {msg.Text}
                    <br />
                    <span className="message-data-time">
                      {new Date(msg.Timestamp).toLocaleString()}
                    </span>
                  </div>
                </li>
              );
            }
          })}
        </ul>
      </div>
      <div className="chat-message clearfix">
        <div className="input-group mb-0">
          <input
            type="text"
            className="form-control"
            placeholder="Search user"
            required
          />
          <button type="submit" className="input-group-text">
            <i className="fa fa-send"></i>
          </button>
        </div>
      </div>
    </form>
  );
};
