<template>
    <div>
        <h2>WebSocket Chat</h2>

        <label for="username">用户名:</label>
        <input v-model="username" id="username" placeholder="请输入用户名" /><br />

        <label for="password">密码:</label>
        <input v-model="password" id="password" type="password" placeholder="请输入密码" /><br />

        <button @click="login">登录</button>
        <p>{{ loginStatus }}</p>

        <div v-if="isLoggedIn">
            <h3>聊天窗口</h3>

            <label for="messageType">消息类型:</label>
            <select v-model="messageType" id="messageType">
                <option value="1">文本</option>
                <option value="2">图片</option>
                <option value="3">视频</option>
            </select><br />

            <label for="toID">发送给 (To ID):</label>
            <input v-model="toID" type="number" id="toID" placeholder="输入目标用户ID" /><br />

            <label for="message">消息内容:</label>
            <textarea v-model="message" id="message" placeholder="输入消息..."></textarea><br />

            <input type="file" ref="fileInput" style="display:none;" /><br />

            <button @click="sendMessage">发送消息</button>
            <p>{{ serverResponse }}</p>
        </div>
    </div>
</template>

<script>
const WebSocket = require('ws');
export default {
    data() {
        return {
            username: '',
            password: '',
            messageType: '1',
            toID: null,
            message: '',
            serverResponse: '',
            loginStatus: '',
            isLoggedIn: false,
            userID: 0,
            token: '',
            ws: null,
        };
    },
    methods: {
        async login() {
            const response = await fetch('http://127.0.0.1:1226/user/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ username: this.username, password: this.password }),
            });
            const data = await response.json();

            if (data.code === 200) {
                this.token = data.result; // 保存返回的 token
                this.loginStatus = '登录成功';
                this.isLoggedIn = true;
                this.connectWebSocket(this.token);
            } else {
                this.loginStatus = '登录失败: ' + data.message;
            }
        },
        connectWebSocket(token) {
            this.ws = new WebSocket("ws://127.0.0.1:1226/ws/chat", {
                "headers": {
                    'Authorization': `Bearer ${token}`
                }
            });

            this.ws.onopen = () => {
                console.log('WebSocket 已连接');
            };

            this.ws.onmessage = (event) => {
                const response = JSON.parse(event.data);
                this.serverResponse = event.data;

                // 解析用户ID
                if (response.type === 'Server' && response.content.includes('已连接至服务器')) {
                    const regex = /用户(\d+)已连接至服务器/;
                    const match = regex.exec(response.content);
                    if (match) {
                        this.userID = parseInt(match[1], 10);
                    }
                }
            };

            this.ws.onclose = () => {
                console.log('WebSocket 已关闭');
            };

            this.ws.onerror = (error) => {
                console.error('WebSocket 错误:', error);
            };
        },
        sendMessage() {
            const type = parseInt(this.messageType);
            const toID = parseInt(this.toID);

            if (type === 2 || type === 3) {
                // 上传图片或视频
                const fileInput = this.$refs.fileInput;
                const file = fileInput.files[0];

                if (!file) {
                    alert('请选择文件');
                    return;
                }

                const reader = new FileReader();
                reader.onload = () => {
                    const fileData = reader.result;

                    // 构造 SendMessage
                    const sendMessage = {
                        type: type,
                        id: this.userID,
                        to_id: toID,
                        data: fileData,
                    };

                    this.ws.send(JSON.stringify(sendMessage));
                };
                reader.readAsDataURL(file);
            } else {
                // 发送文本消息
                const sendMessage = {
                    type: type,
                    id: this.userID,
                    to_id: toID,
                    data: this.message,
                };

                this.ws.send(JSON.stringify(sendMessage));
            }
        },
    },
};
</script>

<style scoped>
/* 添加一些基本样式 */
h2,
h3 {
    margin: 10px 0;
}

label {
    display: block;
    margin: 5px 0;
}
</style>