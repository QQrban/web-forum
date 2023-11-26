export default class Modal {
  constructor(signUpSelector, logInSelector) {
    this.signUpWindow = document.querySelector(signUpSelector);
    this.logInWindow = document.querySelector(logInSelector);
    this.signUpBtn = document.querySelector(".header-btns__signUp");
    this.logInBtn = document.querySelector(".header-btns__logIn");
    this.closeSignUpBtn = document.getElementById("close__signUp");
    this.closeLogInBtn = document.getElementById("close__logIn");
    this.redirectToSignUp = document.getElementById("logIn__window__signUp");
    this.redirectToLogIn = document.getElementById("signUp__window__logIn");
    this.overlay = document.getElementById("overlay");
    this.avatarsContainer = document.getElementById("select-avatar");

    this.addEventListeners();
  }

  addEventListeners() {
    this.signUpBtn.addEventListener("click", (e) => {
      this.showSignUpWindow();
      e.stopPropagation();
    });

    this.logInBtn.addEventListener("click", (e) => {
      this.showLogInWindow();
      e.stopPropagation();
    });

    document.addEventListener("click", (e) => {
      if (
        !this.signUpWindow.contains(e.target) &&
        !this.logInWindow.contains(e.target) &&
        e.target !== this.overlay &&
        e.target !== this.avatarsContainer
      ) {
        this.hideAllWindows();
      }
    });

    this.signUpWindow.addEventListener("click", (e) => e.stopPropagation());
    this.logInWindow.addEventListener("click", (e) => e.stopPropagation());

    this.closeSignUpBtn.addEventListener("click", (e) => {
      this.hideSignUpWindow();
      e.stopPropagation();
    });

    this.closeLogInBtn.addEventListener("click", (e) => {
      this.hideLogInWindow();
      e.stopPropagation();
    });

    this.redirectToSignUp.addEventListener("click", (e) => {
      this.showSignUpWindow();
      e.stopPropagation();
    });

    this.redirectToLogIn.addEventListener("click", (e) => {
      this.showLogInWindow();
      e.stopPropagation();
    });
  }

  hideAllWindows() {
    this.signUpWindow.style.display = "none";
    this.logInWindow.style.display = "none";
  }

  showSignUpWindow() {
    this.hideAllWindows();
    this.signUpWindow.style.display = "block";
  }

  hideSignUpWindow() {
    this.signUpWindow.style.display = "none";
  }

  showLogInWindow() {
    this.hideAllWindows();
    this.logInWindow.style.display = "block";
  }

  hideLogInWindow() {
    this.logInWindow.style.display = "none";
  }
}
