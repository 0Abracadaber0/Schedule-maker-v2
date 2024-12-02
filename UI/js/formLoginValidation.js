const loginForm = document.getElementById('login-form');  // форма входа
const loginUsername = document.getElementById('login-username');
const loginPassword = document.getElementById('login-password');


const users = [
    { username: 'Zhenya', password: 'Password123!' },
    { username: 'Vanya', password: 'Password456!' },
    { username: 'Igor', password: 'Password789!' },
    { username: 'Maks', password: 'Password789!' }
];


const setError = (element, message) => {
    const inputControl = element.parentElement;
    const errorDisplay = inputControl.querySelector('.error');
    errorDisplay.innerText = message;
    inputControl.classList.add('error');
    inputControl.classList.remove('success');
};


const setSuccess = (element) => {
    const inputControl = element.parentElement;
    const errorDisplay = inputControl.querySelector('.error');
    errorDisplay.innerText = '';
    inputControl.classList.add('success');
    inputControl.classList.remove('error');
};

// Валидация для формы входа
const validateLoginInputs = () => {
    const usernameValue = loginUsername.value.trim();
    const passwordValue = loginPassword.value.trim();

    // Очистка ошибок перед новой проверкой
    setSuccess(loginUsername);
    setSuccess(loginPassword);

    // Проверка логина
    if (usernameValue === '') {
        setError(loginUsername, 'Логин обязателен');
        return; 
    }

    // Проверка на наличие пользователя с таким логином
    const user = users.find(user => user.username === usernameValue);

    if (!user) {
        setError(loginUsername, 'Неверный логин');
        return; 
    }

    if (passwordValue === '') {
        setError(loginPassword, 'Пароль обязателен');
        return; 
    }

    if (user.password !== passwordValue) {
        setError(loginPassword, 'Неверный пароль');
        return; 
    }

    // Если логин и пароль верны
    alert('Вход успешен!');
    // перенаправление на другую страницу
    // window.location.href = '/main';
};

// обработчика события для формы входа
loginForm.addEventListener('submit', (e) => {
    e.preventDefault();
    validateLoginInputs(); 
});
