import React, { Component } from "react";
import GlobalUserCard from "../GlobalUserCard";
import { Navigate } from "react-router-dom";
import Cookies from "js-cookie";

export default class Global extends Component {
  constructor(props) {
    super();
    super.constructor(props);

    this.state = {
      authenticated: !(
        Cookies.get("token") === null || Cookies.get("token") === ""
      ),
      users: [],
    };


    const myHeaders = new Headers();
    myHeaders.append("Authorization", `Bearer ${Cookies.get("token")}`);

    if (this.state.authenticated) {
      fetch("http://localhost:8080/svc/home?format=json", {
        method: "GET",
        mode: "cors",
        headers: myHeaders,
      })
        .then((response) => {
          console.log(response.status);
          if (response.status === 401) {
            Cookies.remove("token");
            this.setState({ authenticated: null });
            throw new Error("HTTP status " + response.status);
          }
          if (!response.ok) {
            throw new Error("HTTP status " + response.status);
          }
          return response.json();
        })
        .then((data) => {
          if (data != null) {
            this.setState({ users: data.data });
          }
        })
        .catch((error) => {
          console.log("apiError:", error);
        });
    }
  }

  handleLogout = (event) => {
    event.preventDefault();

    fetch("http://localhost:8080/v1/logout", {
      method: "GET",
      mode: "cors",
    })
      .then((response) => {
        return response.json();
      })
      .then((data) => {
        Cookies.remove("token");
        this.setState({ authenticated: null });
      })
      .catch((error) => {
        console.log("apiError:", error);
      });
  };

  render() {
    if (!this.state.authenticated) {
      return <Navigate to="/sign-in" />;
    }

    return (
      <div>
        <h3 className="font-dancing-script">Find Friends</h3>
        <div className="mb-3">
          <input
            type="text"
            className="form-control"
            placeholder="Search user"
            required
          />
        </div>
        <div className="mb-3 userList">
          {this.state.users.map((user) => {
            return (
              <GlobalUserCard
                person={{
                  user_id: user.user_id,
                  name: user.username,
                  created_at: user.created_at,
                  following: user.following,
                }}
              />
            );
          })}
        </div>

        <form onSubmit={this.handleLogout}>
          <button
            className="btn btn-sm btn-outline-primary w-100 mr-1"
            type="submit"
          >
            Logout
          </button>
        </form>
      </div>
    );
  }
}
