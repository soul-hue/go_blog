
import { createApp } from 'vue'
import {
    ElButton,
    ElForm,
    ElFormItem,
    ElInput,
    ElCard,
    ElContainer,
    ElMessage,
    ElMenu,
    ElMenuItem,
    ElMenuItemGroup,
    ElIcon,
    ElMain,
    ElAside,
    ElHeader,
    ElFooter
} from 'element-plus'
import App from '../App.vue'

const app = createApp(App)
app.use(ElButton)
app.use(ElForm)
app.use(ElFormItem)
app.use(ElInput)
app.use(ElCard)
app.use(ElContainer)
app.use(ElMenu)
app.use(ElMenuItem)
app.use(ElMenuItemGroup)
app.use(ElIcon)
app.use(ElMain)
app.use(ElAside)
app.use(ElHeader)
app.use(ElFooter)
// app._context.config.globalProperties.$message = ElMessage
app.config.globalProperties.$message = ElMessage


export default {}



