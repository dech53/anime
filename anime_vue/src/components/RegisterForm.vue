<template>
    <div class="regist-form">
      <h1>注册</h1>
      <form @submit.prevent="regist">
        <div class="form-group">
          <label for="username">用户名：</label>
          <input type="text" id="username" v-model="username" required/>
        </div>
        <div class="form-group">
          <label for="password">密码：</label>
          <input type="password" id="password" v-model="password" required/>
        </div>
        <div class="form-group">
          <label for="email">邮箱：</label>
          <input type="email" id="email" v-model="email" required/>
        </div>
        <button type="submit">注册</button>
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
        email:'',
        error: ''
      };
    },
    methods: {
      async regist() {
        try {
          const response = await this.$axios.post('http://localhost:1226/regist', {
            username: this.username,
            password: this.password,
            email:this.email,
          });
          if (response.data.code === 200) {
            this.$router.push('/login'); // 登录成功后重定向到首页
          }else{
            this.error = response.data.result
          }
        } catch (error) {
          console.error('注册失败:', error);
        }
      }
  
    }
  };
  </script>
  
  <style scoped>
  .regist-form {
    max-width: 400px;
    margin: 0 auto;
    padding: 30px;
    background-color: #fff;
    border-radius: 8px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  }
  
  h1 {
    font-size: 28px;
    margin-bottom: 30px;
    color: #333;
    text-align: center;
  }
  
  .form-group {
    margin-bottom: 20px;
  }
  
  label {
    display: block;
    font-size: 16px;
    color: #555;
    margin-bottom: 10px;
  }
  
  input[type="text"],
  input[type="password"],input[type="email"] {
    width: calc(100% - 20px);
    padding: 12px;
    border: 1px solid #ddd;
    border-radius: 6px;
    font-size: 16px;
    box-sizing: border-box;
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
  }
  
  button:hover {
    background-color: #0056b3;
  }
  
  .error-message {
    margin-top: 20px;
    color: #d9534f;
    font-size: 14px;
    text-align: center;
  }
  </style>
  