import React from "react";
import "../node_modules/bootstrap/dist/css/bootstrap.min.css";
import "./App.css";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Login from "./Login";
import Global from "./Global";
import SignUp from "./Register";
import Chat from "./Chat";

function App() {
  return (
    <Router>
      <link
          href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css"
          rel="stylesheet"
        />
      <div className="App">
        <div className="auth-wrapper">
          <h1 className="text-white main-title">ChatterBox</h1>
          <div className="auth-inner">
            <Routes>
              <Route exact path="/" element={<Chat />} />
              <Route path="/sign-in" element={<Login />} />
              <Route path="/sign-up" element={<SignUp />} />
              <Route exact path="/home" element={<Global />} />
              <Route exact path="/chat" element={<Chat />} />
            </Routes>
          </div>
        </div>
      </div>
    </Router>
  );
}
export default App;
