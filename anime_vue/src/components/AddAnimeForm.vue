<template>
    <div class="modal" v-if="visible">
        <div class="modal-content">
            <h2>添加 Anime 信息</h2>
            <form @submit.prevent="submitForm">
                <div class="form-group" v-for="(value, key) in animeFields" :key="key">
                    <div class="input-container">
                        <input :type="getInputType(key)" v-model="anime[key]" @focus="focusField(key)" @blur="blurField(key)" required />
                        <label :class="{ active: isFocused[key] || anime[key] }">{{ value }}</label>
                    </div>
                </div>
                <button type="submit">添加</button>
                <button type="button" @click="close">取消</button>
            </form>
        </div>
    </div>
</template>

<script>
export default {
    props: {
        visible: {
            type: Boolean,
            required: true
        },
        anime: {
            type: Object,
            default: () => ({
                Year: '',
                Name: '',
                Season: '',
                EpisodeCount: '',
                Author: '',
                Path: '',
                Cover: ''
            })
        }
    },
    data() {
        return {
            isFocused: {
                Year: false,
                Name: false,
                Season: false,
                EpisodeCount: false,
                Author: false,
                Path: false,
                Cover: false
            },
            animeFields: {
                Year: '年份',
                Name: '名称',
                Season: '季数',
                EpisodeCount: '集数',
                Author: '作者',
                Path: '路径',
                Cover: '封面'
            }
        };
    },
    methods: {
        getInputType(key) {
            return key === 'Year' || key === 'EpisodeCount' ? 'number' : 'text';
        },
        focusField(field) {
            this.isFocused[field] = true;
        },
        blurField(field) {
            this.isFocused[field] = false;
        },
        close() {
            this.$emit('close');
        },
        submitForm() {
            this.$emit('submit', { ...this.anime });
            this.close();
        }
    }
};
</script>

<style scoped>
.modal {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: rgba(0, 0, 0, 0.5);
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 9999;
}

.modal-content {
    background: white;
    padding: 20px;
    border-radius: 8px;
    width: 80%;
    max-width: 600px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.modal-content h2 {
    margin-top: 0;
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
input[type="number"] {
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

input:focus + label,
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

button[type="button"] {
    background-color: #ccc;
    color: black;
}

button[type="button"]:hover {
    background-color: #999;
}
</style>
