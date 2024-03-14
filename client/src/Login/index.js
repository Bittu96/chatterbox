import React, { Component } from "react";
import Cookies from "js-cookie";
import { Navigate } from "react-router-dom";


export default class Login extends Component {
  constructor(props) {
    super();
    super.constructor(props);

    this.state = {
      authenticated: false,
    };
  }

  handleLogin = (event) => {
    event.preventDefault();
    let postData = JSON.stringify({
      username: event.target[0].value,
      password: event.target[1].value,
    });

    fetch("http://localhost:8080/v1/login?format=json", {
      method: "POST",
      mode: "cors",
      body: postData,
    })
      .then((response) => {
        return response.json();
      })
      .then((data) => {
        if (data.data.token !== null && data.data.token !== "") {
          Cookies.set("token", data.data.token, { expires: 1 });
          this.setState({authenticated: true})
        }
      })
      .catch((error) => {
        console.log("apiError:", error);
      });
  };

  render() {
    if( this.state.authenticated) {
      // return <Global />;
      return <Navigate to="/home" />
    }

    return (
      <form onSubmit={this.handleLogin}>
        <h3>Sign In</h3>
        <div className="mb-3">
          <label>Username</label>
          <input
            type="text"
            className="form-control"
            placeholder="Enter Username"
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
        <div className="mb-3">
          <div className="custom-control custom-checkbox">
            <input
              type="checkbox"
              className="custom-control-input"
              id="customCheck1"
            />
            <label className="custom-control-label" htmlFor="customCheck1">
                Remember me
            </label>
          </div>
        </div>
        <div className="d-grid">
          <button type="submit" className="btn btn-primary">
            Submit
          </button>
        </div>
        <p className="forgot-password text-left">
          New User <a href="/sign-up">sign up?</a>
        </p>
        <p className="forgot-password text-right">
          Forgot <a href="#head">password?</a>
        </p>
      </form>
    );
  }
}
