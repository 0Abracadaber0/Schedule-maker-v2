const form = document.getElementById('form');  // форма регистрации
const registerUsername = document.getElementById('register-username');
const registerPassword = document.getElementById('register-password');
const registerPassword2 = document.getElementById('register-password2');

// Функция для обработки ошибок
const setError = (element, message) => {
    const inputControl = element.parentElement;
    const errorDisplay = inputControl.querySelector('.error');
    errorDisplay.innerText = message;
    inputControl.classList.add('error');
    inputControl.classList.remove('success');
};


// Функция для обработки успешной валидации
const setSuccess = (element) => {
    const inputControl = element.parentElement;
    const errorDisplay = inputControl.querySelector('.error');
    errorDisplay.innerText = '';
    inputControl.classList.add('success');
    inputControl.classList.remove('error');
};


// Функция для проверки специальных символов
const containsSpecialCharacter = (str) => {
    const specialCharPattern = /[!@#$%^&*(),.?":{}|<>]/;
    return specialCharPattern.test(str);
};


// Валидация для формы регистрации
const validateRegistrationInputs = () => {
    const usernameValue = registerUsername.value.trim();
    const passwordValue = registerPassword.value.trim();
    const password2Value = registerPassword2.value.trim();

    // Проверка логина
    if (usernameValue === '') {
        setError(registerUsername, 'Логин обязателен');
    } else if (!/^[a-zA-Z0-9._-]{3,}$/.test(usernameValue)) {
        setError(registerUsername, 'Логин должен быть не менее 3 символов и содержать только буквы, цифры, _ или -');
    } else {
        setSuccess(registerUsername);
    }

    // Проверка пароля
    if (passwordValue === '') {
        setError(registerPassword, 'Пароль обязателен');
    } else if (passwordValue.length < 8) {
        setError(registerPassword, 'Пароль должен содержать не менее 8 символов');
    } else if (!/[A-Z]/.test(passwordValue)) {
        setError(registerPassword, 'Пароль должен содержать хотя бы одну заглавную букву');
    } else if (!/\d/.test(passwordValue)) {
        setError(registerPassword, 'Пароль должен содержать хотя бы одну цифру');
    } else if (!containsSpecialCharacter(passwordValue)) {
        setError(registerPassword, 'Пароль должен содержать хотя бы один специальный символ');
    } else {
        setSuccess(registerPassword);
    }

    // Проверка подтверждения пароля
    if (password2Value === '') {
        setError(registerPassword2, 'Подтвердите пароль');
    } else if (password2Value !== passwordValue) {
        setError(registerPassword2, 'Пароли не совпадают');
    } else {
        setSuccess(registerPassword2);
    }
};


// Добавление обработчика события для формы регистрации
form.addEventListener('submit', (e) => {
    e.preventDefault();
    validateRegistrationInputs();  // Валидация для формы регистрации
});
