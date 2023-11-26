export default class PopUpModal {
  constructor(primaryMessage, secondaryMessage, style, time) {
    this.primaryMessageSelector = document.querySelector(
      ".pop-up__modal__text-1"
    );
    this.secondaryMessageSelector = document.querySelector(
      ".pop-up__modal__text-2"
    );
    this.primaryMessage = primaryMessage;
    this.secondaryMessage = secondaryMessage;
    this.closeBtn = document.querySelector("#pop-up__modal__close");
    this.popupModal = document.querySelector(".pop-up__modal");
    this.successSVG = document.querySelector("#pop-up__modal__success-svg");
    this.errorSVG = document.querySelector("#pop-up__modal__error-svg");
    this.time = time;
    this.timerOne = null;
    this.style = style;

    this.addEventListener();
  }

  addEventListener() {
    this.closeBtn.addEventListener("click", () => {
      this.closeModal();
    });
  }

  showModal() {
    if (this.timerOne) {
      clearTimeout(this.timerOne);
    }
    if (this.popupModal.classList.contains("pop-up__modal-active")) {
      return;
    }
    if (this.style === "success") {
      this.errorSVG.classList.add("hidden");
      this.successSVG.classList.remove("hidden");
      this.popupModal.classList.remove("error");
      this.popupModal.classList.add("success");
      this.primaryMessageSelector.classList.remove("error");
      this.primaryMessageSelector.classList.add("success");
    } else if (this.style === "error") {
      this.successSVG.classList.add("hidden");
      this.errorSVG.classList.remove("hidden");
      this.popupModal.classList.remove("success");
      this.popupModal.classList.add("error");
      this.primaryMessageSelector.classList.remove("success");
      this.primaryMessageSelector.classList.add("error");
    }
    this.primaryMessageSelector.innerHTML = this.primaryMessage;
    this.secondaryMessageSelector.innerHTML = this.secondaryMessage;
    this.popupModal.style.display = "block";
    this.popupModal.classList.add("pop-up__modal-active");
    this.timerOne = setTimeout(() => {
      this.popupModal.classList.remove("pop-up__modal-active");
    }, this.time);
  }

  closeModal() {
    this.popupModal.classList.remove("pop-up__modal-active");
    this.popupModal.style.display = "none";
    clearTimeout(this.timerOne);
  }
}
