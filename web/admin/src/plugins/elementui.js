
import { createApp } from 'vue'
import {
    ElButton,
    ElForm,
    ElFormItem,
    ElInput,
    ElCard,
    ElContainer,
    ElMessage
} from 'element-plus'
import App from '../App.vue'

const app = createApp(App)
app.use(ElButton)
app.use(ElForm)
app.use(ElFormItem)
app.use(ElInput)
app.use(ElCard)
app.use(ElContainer)
// app._context.config.globalProperties.$message = ElMessage
app.config.globalProperties.$message = ElMessage


export default {}



