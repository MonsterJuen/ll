<template>
  <div class="inbound-form">
    <h2>商品入库</h2>
    <el-form :model="form" :rules="rules" ref="form" label-width="120px">
      <el-form-item label="商品名称" prop="productName">
        <el-input v-model="form.productName" placeholder="请输入商品名称"></el-input>
      </el-form-item>
      
      <el-form-item label="数量" prop="quantity">
        <el-input-number v-model="form.quantity" :min="1" placeholder="请输入数量"></el-input-number>
      </el-form-item>
      
      <el-form-item label="单位" prop="unit">
        <el-select v-model="form.unit" placeholder="请选择单位">
          <el-option label="个" value="个"></el-option>
          <el-option label="箱" value="箱"></el-option>
          <el-option label="千克" value="千克"></el-option>
        </el-select>
      </el-form-item>
      
      <el-form-item label="温度" prop="temperature">
        <el-input-number 
          v-model="form.temperature" 
          :precision="1" 
          :step="0.1" 
          :min="-50"
          :max="50"
          placeholder="请输入温度">
          <template slot="append">°C</template>
        </el-input-number>
      </el-form-item>
      
      <el-form-item label="存放位置" prop="location">
        <el-input v-model="form.location" placeholder="请输入存放位置"></el-input>
      </el-form-item>
      
      <el-form-item>
        <el-button 
          type="primary" 
          @click="submitForm" 
          :loading="loading"
          :disabled="loading">
          {{ loading ? '提交中...' : '提交' }}
        </el-button>
        <el-button @click="resetForm" :disabled="loading">重置</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'InboundForm',
  data() {
    return {
      form: {
        productName: '',
        quantity: 1,
        unit: '',
        temperature: 0,
        location: ''
      },
      rules: {
        productName: [
          { required: true, message: '请输入商品名称', trigger: 'blur' },
          { min: 2, max: 50, message: '长度在 2 到 50 个字符', trigger: 'blur' }
        ],
        quantity: [
          { required: true, message: '请输入数量', trigger: 'blur' }
        ],
        unit: [
          { required: true, message: '请选择单位', trigger: 'change' }
        ],
        temperature: [
          { required: true, message: '请输入温度', trigger: 'blur' }
        ],
        location: [
          { required: true, message: '请输入存放位置', trigger: 'blur' },
          { min: 2, max: 100, message: '长度在 2 到 100 个字符', trigger: 'blur' }
        ]
      },
      loading: false
    }
  },
  methods: {
    submitForm() {
      this.$refs.form.validate(async (valid) => {
        if (valid) {
          this.loading = true
          try {
            const response = await axios.post('http://localhost:8080/api/inbound', this.form)
            this.$message.success(response.data.message)
            this.resetForm()
          } catch (error) {
            if (error.response) {
              // 服务器返回错误
              this.$message.error(error.response.data.error || '入库失败')
              if (error.response.data.details) {
                console.error('错误详情:', error.response.data.details)
              }
            } else if (error.request) {
              // 请求未收到响应
              this.$message.error('无法连接到服务器，请检查网络连接')
            } else {
              // 请求配置出错
              this.$message.error('请求配置错误：' + error.message)
            }
          } finally {
            this.loading = false
          }
        }
      })
    },
    resetForm() {
      this.$refs.form.resetFields()
    }
  }
}
</script>

<style scoped>
.inbound-form {
  max-width: 600px;
  margin: 20px auto;
  padding: 20px;
  border: 1px solid #dcdfe6;
  border-radius: 4px;
}
</style> 