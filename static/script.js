document.addEventListener("DOMContentLoaded", function() {
  const loginText = document.querySelector(".title-text .login");
  const loginForm = document.querySelector("form.login");
  const signupForm = document.querySelector("form.signup");
  const loginBtn = document.querySelector("label.login");
  const signupBtn = document.querySelector("label.signup");
  const signupLink = document.querySelector("form .signup-link a");

  signupBtn.onclick = (()=>{
    loginForm.classList.remove('active');
    signupForm.classList.add('active');
    loginText.style.marginLeft = "-50%";
  });

  loginBtn.onclick = (()=>{
    loginForm.classList.add('active');
    signupForm.classList.remove('active');
    loginText.style.marginLeft = "0%";
  });

  signupLink.onclick = (()=>{
    signupBtn.click();
    return false;
  });
});
