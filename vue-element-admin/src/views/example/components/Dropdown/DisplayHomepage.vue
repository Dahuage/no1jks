<template>
  <el-dropdown :show-timeout="100" trigger="click">
    <el-button plain>
      {{ !display_homepage?'不展示在首页':'展示在首页' }}
      <i class="el-icon-caret-bottom el-icon--right" />
    </el-button>
    <el-dropdown-menu slot="dropdown" class="no-padding">
      <el-dropdown-item>
        <el-radio-group v-model="display_homepage" style="padding: 10px;">
          <el-radio :label="1">
            首页展示
          </el-radio>
          <el-radio :label="0">
            不展示
          </el-radio>
        </el-radio-group>
      </el-dropdown-item>
    </el-dropdown-menu>
  </el-dropdown>
</template>

<script>
export default {
  props: {
    value: {
      type: Number,
      default: 1
    }
  },
  computed: {
    display_homepage: {
      get() {
        return this.value
      },
      set(val) {
        this.$emit('input', val)
      }
    }
  }
}
</script>
