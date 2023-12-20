<script>
import {reactive, ref, watch} from 'vue'
import {CalFee} from '../../wailsjs/go/main/App'
export default {
  setup() {
    const form = reactive({
      goods: 0,
      goodsRes: 0,
      service: 0,
      serviceRes: 0,
      engin: 0,
      enginRes: 0,
    })


// function greet() {
//   // Greet(data.name).then(result => {
//   //   data.resultText = result
//   // })
//   CalFee(data.name, 2).then(res => {
//     data.resultText = res
//   })
// }
//

    const handleChange = (amount, type) => {
      CalFee(amount, type).then(res => {
        if (type === 0){
          form.goodsRes = res
        } else if (type === 1){
          form.serviceRes = res
        } else if (type === 2){
          form.enginRes = res
        }
      })
    }

    const ruleForm = ref(null)

    const resetForm = () => {
      ruleForm.value.resetFields()
    }

    watch(() => {
      handleChange(form.goods, 0)
      handleChange(form.engin, 2)
      handleChange(form.service, 1)
    }, { immediate: true })

    return { form, resetForm, ruleForm }
  }
}


</script>

<template>
  <main>
    <div id="result" class="result">服务费计算器</div>
    <el-form :model="form" label-width="120px" class="form" ref="ruleForm">
      <el-form-item label="货物：" prop="goods">
        <el-input-number controls-position="" class="input-box"  v-model="form.goods" :min="0" />
        <el-input-number controls-position="" class="input-box" disabled=true	 v-model="form.goodsRes" :min="0" />
      </el-form-item>
      <el-form-item label="服务：" prop="service">
        <el-input-number controls-position="" class="input-box" 	 v-model="form.service" :min="0" />
        <el-input-number controls-position="" class="input-box" disabled=true	 v-model="form.serviceRes" :min="0" />
      </el-form-item>
      <el-form-item label="工程：" prop="engin">
        <el-input-number controls-position="" class="input-box" 	 v-model="form.engin" :min="0" />
        <el-input-number controls-position="" class="input-box" disabled=true	 v-model="form.enginRes" :min="0" />
      </el-form-item>
      <el-form-item>
        <el-button @click="resetForm">Reset</el-button>
      </el-form-item>
    </el-form>
  </main>
</template>

<style scoped>
.form {
  width: 70%;
  margin: 1.5rem auto;
}
.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
  color: #1b2636;
}

.input-box{
  width: 200px;
  height: 30px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 0 0 0 20px;
  padding: 0 8px;
  cursor: pointer;
}

.input-box .btn:hover {
  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  color: #333333;
}

.input-box .input {
  border: none;
  border-radius: 3px;
  outline: none;
  height: 30px;
  line-height: 30px;
  padding: 0 10px;
  background-color: rgba(240, 240, 240, 1);
  -webkit-font-smoothing: antialiased;
}

.input-box .input:hover {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}
</style>
