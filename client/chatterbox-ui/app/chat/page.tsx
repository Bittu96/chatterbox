"use client";

import Cookies from "js-cookie";
import React, { useRef, useEffect, useState } from "react";

export default function Chat() {
  console.log("checking session");

  let username = Cookies.get("session-username")

  return (
    <div className="flex flex-col h-screen w-screen items-center justify-center">
      <h3 className="title text-5xl text-white mb-2">Chatterbox</h3>
      <h3>welcome {username}</h3>
      <div className="z-10 w-full max-w-md overflow-hidden rounded-2xl border border-gray-100 shadow-xl bg-white">
        <div className="flex flex-col items-center justify-center space-y-3 bg-white px-4 py-4 pt-8 text-center sm:px-16"></div>

        <TextBx senderId={username} receiverId={123} token={"algnsegks"} />

        <div className="items-center justify-center  border-b border-gray-200 bg-white py-6 pt-8 text-center"></div>
      </div>
    </div>
  );
}

const useWs = (senderId: any, receiverId: any, token: any) => {
  console.log("useWs:", senderId);

  const [isReady, setIsReady] = useState(false);
  const [val, setVal] = useState(null);
  // const ws = useRef(null);

  const socket = new WebSocket(
    `ws://localhost:8080/webs/chat/${senderId}/${receiverId}/${token}`
  );

  useEffect(() => {
    socket.onopen = () => setIsReady(true);
    socket.onclose = () => setIsReady(false);
    socket.onmessage = (event: any) => setVal(event.data);

    // ws.current = socket;
    // socket.send.bind(ws)

    return () => {
      // socket.close();
    };
  }, []);

  // bind is needed to make sure `send` references correct `this`
  return [isReady, val, socket.send.bind(socket)];
};

const TextBx = ({ senderId, receiverId, token }: any) => {
  console.log("textbox:", senderId);
  const [ready, val, send] = useWs(senderId, receiverId, token);

  let msgs = [];
  let dataFound = val != null;
  if (dataFound) {
    msgs = JSON.parse(val);
    console.log(msgs);
  }

  var handleSubmit = (event: any) => {
    event.preventDefault();
    send(event.target[0].value);
  };

  return (
    <form onSubmit={handleSubmit} className="chat flex flex-col items-center">
      <div className="chat-history">
        <ul className="">
          {msgs.map((msg: any) => {
            console.log(msg);
            if (senderId != msg.UserId) {
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
      <div className="chat-message border-red-400 w-100">
        <input
          type="text"
          className="form-control w-100"
          placeholder="write you message.."
          required
        />
      </div>
    </form>
  );
};
