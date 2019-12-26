<template>
  <el-dropdown :show-timeout="100" trigger="click">
    <el-button plain>
      {{ !display_homepage?'首页展示':'首页不展示' }}
      <i class="el-icon-caret-bottom el-icon--right" />
    </el-button>
    <el-dropdown-menu slot="dropdown" class="no-padding">
      <el-dropdown-item>
        <el-radio-group v-model="display_homepage" style="padding: 10px;">
          <el-radio :label="true">
            首页展示
          </el-radio>
          <el-radio :label="false">
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
      type: Boolean,
      default: false
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
