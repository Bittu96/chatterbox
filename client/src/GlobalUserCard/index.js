import React, { Component } from "react";
import Cookies from "js-cookie";

export default class GlobalUserCard extends Component {
  constructor(props) {
    super();
    super.constructor(props);

    var { id, name, created_at, following } = this.props.person;
    this.name = name;
    this.created_at = created_at;
    this.id = id;
    this.following = following;
    
    var random = Math.floor((Math.random() * 100) + 1)
    this.pic_link = `https://picsum.photos/id/${random}/200/200`
    this.state = {
      following: !(this.following === -1),
    };
  }

  prettifyDateTime = (str) => {
    const [date, time] = str.split("T");
    const [year, month, day] = date.split("-");
    return `${day}/${month}/${year}`;
  };

  handleFollow = (event) => {
    event.preventDefault();

    const myHeaders = new Headers();
    myHeaders.append("Authorization", `Bearer ${Cookies.get("token")}`);

    fetch(`http://localhost:8080/v2/follow?following_id=${this.id}`, {
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
        if (data.status === "success") {
          this.setState({ following: true });
        }
      })
      .catch((error) => {
        console.log("apiError:", error);
      });
  };

  handleUnFollow = (event) => {
    event.preventDefault();

    const myHeaders = new Headers();
    myHeaders.append("Authorization", `Bearer ${Cookies.get("token")}`);

    fetch(`http://localhost:8080/v2/unfollow?following_id=${this.id}`, {
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
          if (data.status === "success") {
            this.setState({ following: false });
          }
        }
      })
      .catch((error) => {
        console.log("apiError:", error);
      });
  };

  render() {
    return (
      <div className="mt-2">
        <div className="container d-flex justify-content-center">
          <div className="global-user-card user-card p-2">
            <div className="d-flex align-items-center">
              <div className="image">
                <img
                  alt="head"
                  src={this.pic_link}
                  className="rounded"
                  width="100"
                />
              </div>

              <div className="m-2 w-100">
                <h6 className="mb-0 mt-0">{this.name}</h6>
                <small className="xsmall">
                  member since {this.prettifyDateTime(this.created_at)}
                </small>

                {/* <div className="p-2 bg-primary d-flex justify-content-between rounded text-white stats">
                  <div className="d-flex flex-column m-2s">
                    <span className="followers">Followers</span>
                    <span className="number2">980</span>
                  </div>

                  <div className="d-flex flex-column">
                    <span className="followers">Following</span>
                    <span className="number2">700</span>
                  </div>
                </div> */}

                <div className="button mt-2 d-flex flex-row align-items-center">
                  <button className="btn btn-sm btn-outline-primary w-100 mr-1">
                    Chat
                  </button>
                  <div className="m-1"></div>
                  <button className={this.state.following ? "btn btn-sm btn-secondary w-100 ml-1":"btn btn-sm btn-primary w-100 ml-1"}
                    onClick={
                      this.state.following
                        ? this.handleUnFollow
                        : this.handleFollow
                    }
                  >
                    {this.state.following ? "following" : "follow"}
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    );
  }
}
