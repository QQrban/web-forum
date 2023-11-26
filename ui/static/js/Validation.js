export default class Validation {
  constructor(selector, regex, message) {
    this.regex = regex;
    this.selector = document.getElementById(selector);
    this.message = message;
    this.error = document.getElementById("register__window__error");
    this.isValid = false;

    this.checkInput = this.checkInput.bind(this);

    if (this.selector) {
      this.selector.addEventListener("change", this.checkInput);  
    }
  }

  checkInput(event) {
    const inputValue = event.target.value;
    if (!this.validate(inputValue)) {
      this.error.innerHTML = this.message;
      this.isValid = false;
    } else {
      this.isValid = true;
      this.error.innerHTML = "";
    }
  }

  validate(value) {
    return this.regex.test(value);
  }
}
