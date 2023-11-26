import Navigation from "./Navigation.js";
import Accordion from "./Accordion.js";
import Modal from "./Modal.js";
import Validation from "./Validation.js";
import DropdownMenu from "./DropdownMenu.js";
import { avatars } from "./avatars.js";
import ShowLabel from "./ShowLabel.js";
import Existence from "./Existence.js";
import PopUpModal from "./PopUpModal.js";
import AuthUI from "./AuthUI.js";
import Link from "./Link.js";
import DataFetcher from "./DataFetcher.js";
import { postsTemplate, commentariesTemplate } from "./templates.js";
import Filters from "./Filters.js";
import { hideLoader, showLoader } from "./helpers.js";
import Like from "./Like.js";
import Search from "./Search.js";

const authUi = new AuthUI();
document.addEventListener("DOMContentLoaded", () => {
  authUi.checkSessionStatus();
  const topicLinks = new Link(".panel__item", "/topic/", "not__link");
  const search = new Search();
});

window.addEventListener("pageshow", function (event) {
  var historyTraversal =
    event.persisted ||
    (typeof window.performance != "undefined" &&
      window.performance.navigation.type === 2);
  if (historyTraversal) {
    window.location.reload();
  }
});

if (
  window.location.pathname === "/" ||
  window.location.href.includes("/topic") ||
  window.location.href.includes("/post") ||
  window.location.href.includes("/activity")
) {
  const navigation = new Navigation([
    ".header-logo__container",
    ".welcome-banner",
    ".search-bar__btns-about-home",
  ]);
}

if (window.location.pathname === "/") {
  document.addEventListener("DOMContentLoaded", () => {
    showLoader();
    const topicLinks = new Link(".commentator__topic", "/post/");
    const postName = document.querySelectorAll(
      ".main-page__commentator__topic"
    );
    postName.forEach((post) => {
      if (post.innerHTML.trim().length > 25) {
        post.innerHTML = post.innerHTML.trim().slice(0, 26) + "...";
      }
    });
    setTimeout(() => {
      hideLoader();
    }, 200);
  });
}

if (window.location.href.includes("/topic")) {
  const dropdownMenu = new DropdownMenu(
    "#clickable-btn__category",
    "#filters__categories",
    ".dropdown-svg"
  );
  const showLabel = new ShowLabel(
    "#topic__header__btn-svg",
    ".topic__header__btn-label",
    1000
  );
  document.addEventListener("DOMContentLoaded", function () {
    const selectedTopic = new DataFetcher(
      "/api/topic/",
      ".all-posts__container",
      postsTemplate
    );
    selectedTopic.fetchByID("descending").then(() => {
      document
        .querySelector(".all-posts__container")
        .classList.remove("loading");
      const topicLinks = new Link(".all-posts__item", "/post/", "not__link");
      const reactions = new Like(".likes-dislikes", "post", "notActivity");
    });
    const sortByDate = new Filters();
  });
}

if (
  window.location.href.includes("/post") ||
  window.location.href.includes("/topic")
) {
  const backButton = document.querySelector(".back-button");
  if (backButton) {
    backButton.addEventListener("click", function () {
      window.history.back();
    });
  }
}

if (window.location.href.includes("/activity")) {
  const topicLinks = new Link(".all-posts__item", "/post/", "not__link");
  const reactions = new Like(".likes-dislikes", "post");
}

if (window.location.href.includes("/post")) {
  const selectedPost = new DataFetcher(
    "/api/post/",
    ".post-page__comment__container",
    commentariesTemplate
  );
  selectedPost.fetchByID("ascending").then(() => {
    const reactions = new Like(".likes-dislikes", "comment", "notActivity");
  });
}

const accordion = new Accordion("forums-list__accordion");

const modal = new Modal(".signUp__window", ".logIn__window");

const emailRegex = /^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,6}$/;
const usernameRegex = /^[0-9A-Za-z]{6,16}$/;
const passRegex = /^(?=.*?[0-9])(?=.*?[A-Za-z]).{8,32}$/;
const password = document.getElementById("#signUp__form__password");
const confirmPassword = document.getElementById(
  "#signUp__form__confirmPassword"
);

const checkEmail = new Validation(
  "signUp__form__email",
  emailRegex,
  "Please enter a valid email"
);

const checkEmailExistence = new Existence(
  "signUp__form__email",
  "email",
  "email__validation__error"
);

const checkUsername = new Validation(
  "signUp__form__username",
  usernameRegex,
  "Username must be between 6 to 16 characters long and can only contain letters and numbers"
);

const checkUsernameExistence = new Existence(
  "signUp__form__username",
  "username",
  "username__validation__error"
);

const checkPassword = new Validation(
  "signUp__form__password",
  passRegex,
  "Your password must be between 8 to 32 characters long, contain at least one letter and one number."
);

const avatarsContainer = document.getElementById("select-avatar");
const overlay = document.getElementById("overlay");

overlay.addEventListener("click", () => {
  avatarsContainer.style.display = "none";
  overlay.style.display = "none";
});

document.addEventListener("keydown", function (event) {
  if (event.key === "Escape") {
    const signModal = document.querySelector(".signUp__window");
    const logInModal = document.querySelector(".logIn__window");
    const newPostModal = document.querySelector(".new-post__modal");
    const activityModal = document.querySelector("#remove-like-modal");
    if (signModal) signModal.style.display = "none";
    if (logInModal) logInModal.style.display = "none";
    if (avatarsContainer) avatarsContainer.style.display = "none";
    if (overlay) overlay.style.display = "none";
    if (newPostModal) newPostModal.style.display = "none";
    if (activityModal) activityModal.classList.add("hidden");
  }
});

avatars.forEach((avatar) => {
  avatarsContainer.insertAdjacentHTML(
    "beforeend",
    `
    <img data-img="${avatar}" class="clicked-avatar" src="${avatar}" alt="avatar"/>
  `
  );
});

const clickedAvatar = document.querySelectorAll(".clicked-avatar");
const chosenAvatarImg = document.getElementById("chosen-avatar-img");
const chosenAvatarRadio = document.getElementById("chooseAvatar__radio");

clickedAvatar.forEach((avatar) => {
  avatar.addEventListener("click", (e) => {
    const link = e.target.getAttribute("data-img");
    chosenAvatarImg.setAttribute("src", link);
    chosenAvatarRadio.setAttribute("value", link);
    avatarsContainer.style.display = "none";
    overlay.style.display = "none";
    e.stopPropagation();
  });
});

chosenAvatarImg.addEventListener("click", () => {
  overlay.style.display = "block";
  avatarsContainer.style.display = "grid";
});

document
  .getElementById("signUp__form")
  .addEventListener("submit", async (e) => {
    e.preventDefault();
    const error = document.getElementById("register__window__error");
    const password = document.getElementById("signUp__form__password").value;
    const confirmPassword = document.getElementById(
      "signUp__form__confirmPassword"
    ).value;

    if (
      checkEmail.isValid &&
      checkEmailExistence.isValid &&
      checkUsername.isValid &&
      checkUsernameExistence.isValid &&
      checkPassword.isValid &&
      password === confirmPassword
    ) {
      document.getElementById("register__window__error").innerHTML = "";
      const formData = new FormData(e.target);
      showLoader();
      try {
        const response = await fetch("/register", {
          method: "POST",
          headers: {
            "Content-Type": "application/x-www-form-urlencoded",
          },
          body: new URLSearchParams(formData),
          credentials: "include",
        });
        const data = await response.json();
        if (data.Ok) {
          const isAuthenticated = await authUi.checkSessionStatus();
          if (isAuthenticated) {
            authUi.toggleUI(true);
          }
          const successModal = new PopUpModal(
            "Success!",
            "Your account has been created!",
            "success",
            3500
          );
          successModal.showModal();
          if (window.location.href.includes("/post")) {
            window.location.reload();
          }
        }
      } catch (error) {
        console.error("Registration failed:", await response.text());
      } finally {
        hideLoader();
      }
    } else {
      if (!checkEmail.isValid) {
        error.innerHTML = "Please enter a valid email";
      } else if (!checkUsername.isValid) {
        error.innerHTML =
          "Username must be between 6 to 16 characters long and can only contain letters and numbers";
      }
      if (password.length < 8) {
        error.innerHTML = "Passwords length must be at least 8 symbols";
      } else if (password !== confirmPassword) {
        error.innerHTML = "Passwords do not match";
      }
    }
  });

document.getElementById("logIn__form").addEventListener("submit", async (e) => {
  e.preventDefault();
  const username = e.target[0].value;
  const password = e.target[1].value;
  showLoader();
  try {
    const response = await fetch("/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ username, password }),
      credentials: "include",
    });
    const data = await response.json();
    if (data.Ok) {
      const isAuthenticated = await authUi.checkSessionStatus();
      if (isAuthenticated) {
        authUi.toggleUI(true);
      }
      const successModal = new PopUpModal(
        "Success!",
        "You successfully logged in",
        "success",
        3500
      );
      successModal.showModal();
      if (window.location.href.includes("/post")) {
        window.location.reload();
      }
    } else {
      document.getElementById("login__window__error").innerHTML = data.Message;
    }
  } catch (error) {
    console.error(error);
  } finally {
    hideLoader();
  }
});
