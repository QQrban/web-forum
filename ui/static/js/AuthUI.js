import Comment from "./Comment.js";
import CreatePost from "./CreatePost.js";
import ShowLabel from "./ShowLabel.js";

export default class AuthUI {
  constructor() {
    this.cacheDomElements();
    this.bindMethods();
    this.addEventListeners();
    this.checkPage();
  }

  cacheDomElements() {
    this.registrationBtns = document.querySelector(".header-btns__container");
    this.exitBtn = document.querySelector(".header__log-out__container");
    this.createPostBtn = document.querySelector("#topic__header__btn-svg");
    this.addCommentBtn = document.querySelector("#submit-comment");
    this.commentsLoggedIn = document.querySelector("#active-commentator");
    this.commentsLoggedOut = document.querySelector("#comment__body__guest");
    this.logOutBtn = document.querySelector("#log-out__svg");
  }

  bindMethods() {
    this.boundLogout = this.logout.bind(this);
  }

  addEventListeners() {
    if (this.logOutBtn) {
      this.logOutBtn.addEventListener("click", this.boundLogout);
    }
  }

  checkPage() {
    const isTopicsPage = window.location.href.includes("/topic");
    const isPostPage = window.location.href.includes("/post");
    if (isTopicsPage) {
      this.createPostModal = new CreatePost(false);
    }
    if (isPostPage && this.commentsLoggedIn && this.commentsLoggedOut) {
      this.toggleCommentsVisibility(false);
      this.formAccess = new Comment(false);
    }
  }

  toggleCommentsVisibility(isLoggedIn) {
    if (!this.commentsLoggedIn || !this.commentsLoggedOut) {
      return;
    }
    this.commentsLoggedOut.classList.toggle("hidden", isLoggedIn);
    this.commentsLoggedIn.classList.toggle("hidden", !isLoggedIn);
  }

  toggleUI(isLoggedIn) {
    this.initShowLabel();
    this.registrationBtns.classList.toggle("hidden", isLoggedIn);
    this.exitBtn.classList.toggle("hidden", !isLoggedIn);
    if (this.createPostModal) {
      this.createPostModal.isLogged = isLoggedIn;
      this.createPostModal.updateEventListeners();
    }
    if (this.addCommentBtn) {
      this.formAccess.isLogged = isLoggedIn;
      if (isLoggedIn) {
        this.formAccess.addComment();
      } else {
        this.formAccess.removeComment();
      }
    }
    this.toggleCommentsVisibility(isLoggedIn);
  }

  async checkSessionStatus() {
    try {
      const response = await fetch("/check-session", {
        credentials: "include",
      });
      const data = await response.json();
      if (data.Ok) {
        console.log("Member Session is active");
        this.toggleUI(true);
      } else {
        console.log("Guest Session is active");
        this.toggleUI(false);
      }
    } catch (error) {
      console.error("Session Status error: ", error);
    }
  }

  async logout() {
    try {
      const response = await fetch("/logout", {
        credentials: "include",
      });
      const data = await response.json();
      if (data.Ok) {
        if (window.location.href.includes("/activity")) {
          window.location.href = "/";
        } else {
          window.location.reload();
        }
        this.toggleUI(false);
      } else {
        console.log("Error logging out");
      }
    } catch (error) {
      console.log(error);
    }
  }

  initShowLabel() {
    this.logOutLabel = new ShowLabel(
      ".header__log-out__container",
      ".office__container",
      0
    );
  }
}
