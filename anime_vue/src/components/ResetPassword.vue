<template>
    <div class="login-form">
        <h1>重置密码</h1>
        <form @submit.prevent="reset">
            <div class="form-group">
                <div class="input-container">
                    <input type="text" id="username" v-model="username" @focus="focusField('username')"
                        @blur="blurField('username')" required />
                    <label :class="{ active: isFocused.username || username }">用户名</label>
                </div>
            </div>
            <div class="form-group">
                <div class="input-container">
                    <input type="password" id="vaildatePassword" v-model="vaildatePassword"
                        @focus="focusField('vaildatePassword')" @blur="blurField('vaildatePassword')" required />
                    <label :class="{ active: isFocused.vaildatePassword || vaildatePassword }">验证码</label>
                </div>
            </div>
            <div class="form-group">
                <div class="input-container">
                    <input type="password" id="new_password_one" v-model="new_password_one"
                        @focus="focusField('new_password_one')" @blur="blurField('new_password_one')" required />
                    <label :class="{ active: isFocused.new_password_one || new_password_one }">新密码</label>
                </div>
            </div>
            <div class="form-group">
                <div class="input-container">
                    <input type="password" id="new_password_twice" v-model="new_password_twice" @focus="focusField('new_password_twice')"
                        @blur="blurField('new_password_twice')" required />
                    <label :class="{ active: isFocused.new_password_twice || new_password_twice }">重复新密码</label>
                </div>
            </div>
            <button type="submit">重置密码</button>
        </form>
        <div v-if="error" class="error-message">{{ error }}</div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            username: '',
            password: '',
            new_password_one: '',
            new_password_twice: '',
            error: '',
            isFocused: {
                username: false,
                new_password_one: false,
                new_password_twice: false,
                vaildatePassword: false
            }
        };
    },
    methods: {
        focusField(field) {
            this.isFocused[field] = true;
        },
        blurField(field) {
            this.isFocused[field] = false;
        },
        async reset() {
            if (new_password_one != new_password_twice){
                this.error = "两次密码不一致";
            }
        }
    }
};
</script>

<style scoped>
.login-form {
    max-width: 400px;
    margin: 0 auto;
    padding: 30px;
    background-color: #fff;
    border-radius: 8px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
    display: flex;
    flex-direction: column;
}

h1 {
    font-size: 28px;
    margin-bottom: 30px;
    color: #333;
    text-align: center;
}

.form-group {
    margin-bottom: 20px;
    position: relative;
}

.input-container {
    position: relative;
}

input[type="text"],
input[type="password"] {
    width: 100%;
    padding: 12px;
    border: 1px solid #ddd;
    border-radius: 6px;
    font-size: 16px;
    box-sizing: border-box;
    outline: none;
    transition: border-color 0.2s ease;
}

input:focus {
    border-color: #007bff;
}

label {
    position: absolute;
    left: 12px;
    top: 50%;
    transform: translateY(-50%);
    font-size: 16px;
    color: #aaa;
    transition: 0.2s ease all;
    pointer-events: none;
}

input:focus+label,
label.active {
    top: -10px;
    left: -3px;
    font-size: 12px;
    color: #007bff;
    background: #fff;
    padding: 0 5px;
}

button {
    width: 100%;
    padding: 12px;
    background-color: #007bff;
    border: none;
    border-radius: 6px;
    color: #fff;
    font-size: 16px;
    cursor: pointer;
    transition: background-color 0.2s ease;
    margin-top: 20px;
}

button:hover {
    background-color: #0056b3;
}


a {
    text-decoration: none;
    color: #9aa7b1;
}

a:hover {
    color: #0078d4;
}

.error-message {
    margin-top: 20px;
    color: #d9534f;
    font-size: 14px;
    text-align: center;
}
</style>