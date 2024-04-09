import React, { Component } from "react";
import { Navigate } from "react-router-dom";
import "./../App.css";

export default class SignUp extends Component {
  constructor(props) {
    super();
    super.constructor(props);

    this.state = {
      registered: false,
    };
  }

  handleRegister = (event) => {
    event.preventDefault();
    let postData = JSON.stringify({
      username: event.target[0].value,
      first_name: event.target[1].value,
      last_name: event.target[2].value,
      email: event.target[3].value,
      password: event.target[4].value,
    });

    fetch("http://localhost:8080/auth/register?format=json", {
      method: "POST",
      mode: "cors",
      body: postData,
    })
      .then((response) => {
        return response.json();
      })
      .then((data) => {
        if (data.status === "success") {
          this.setState({
            registered: true,
          });
          return;
        }
      })
      .catch((error) => {
        console.log("apiError:", error);
      });
  };

  render() {
    if (this.state.registered) {
      return <Navigate to="/sign-in" />;
    }

    return (
      <form onSubmit={this.handleRegister}>
        <h3>Sign Up</h3>
        <div className="mb-3">
          <label>Username</label>
          <input
            type="text"
            className="form-control"
            placeholder="username"
            required
          />
        </div>
        <div className="mb-3">
          <label>First name</label>
          <input
            type="text"
            className="form-control"
            placeholder="First name"
          />
        </div>
        <div className="mb-3">
          <label>Last name</label>
          <input type="text" className="form-control" placeholder="Last name" />
        </div>
        <div className="mb-3">
          <label>Email address</label>
          <input
            type="email"
            className="form-control"
            placeholder="Enter email"
            required
          />
        </div>
        <div className="mb-3">
          <label>Password</label>
          <input
            type="password"
            className="form-control"
            placeholder="Enter password"
            required
          />
        </div>
        <div className="d-grid">
          <button type="submit" className="btn btn-primary">
            Sign Up
          </button>
        </div>
        <p className="forgot-password text-right">
          Already registered <a href="/sign-in">sign in?</a>
        </p>
      </form>
    );
  }
}
