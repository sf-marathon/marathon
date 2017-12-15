<template>
  <div class="jh-join">
    <yd-cell-group>
      <yd-cell-item>
        <span slot="left" class="jh-join-title__address">寄件地址</span>
      </yd-cell-item>
      <yd-cell-item>
        <span slot="left">所在地区</span>
        <input slot="right" 
          type="text" 
          ref="form-address"
          style="text-align:right;" 
          @click.stop="showPicker = true" 
          v-model="formData['address']" 
          readonly 
          placeholder="请选择寄件地址">
      </yd-cell-item>
      <yd-cell-item>
        <yd-textarea 
          slot="right" 
          min="5"
          ref="form-detail"
          v-model="formData['detail']" 
          placeholder="请填写详细地址，不少于5个字">
        </yd-textarea>
      </yd-cell-item>
    </yd-cell-group>
    <yd-cell-group>
      <yd-cell-item>
        <span slot="left" class="jh-join-title__people">寄件人</span>
      </yd-cell-item>
      <yd-cell-item>
        <yd-input
          slot="right"
          ref="form-name"
          v-model="formData['name']" 
          required :show-clear-icon="false" 
          placeholder="请填写寄件人姓名">
        </yd-input>        
      </yd-cell-item>
      <yd-cell-item>
        <yd-input 
          slot="right" 
          ref="form-phone"
          v-model="formData['phone']" 
          type="tel" regex="mobile" 
          required :show-clear-icon="false" 
          placeholder="请填写寄件人电话">
        </yd-input>        
      </yd-cell-item>
    </yd-cell-group>
    <yd-cell-group>
      <yd-cell-item>
        <span slot="left" class="jh-join-title__number">预估每日件量<em>（请如实填报，以便于我们为您准备资源）</em></span>
      </yd-cell-item>
      <yd-cell-item>
        <yd-input slot="right" ref="form-number" v-model="formData['number']" type="number" required :show-clear-icon="false" placeholder="请填写寄件数量"></yd-input>
      </yd-cell-item>
      <yd-cell-item arrow type="label">
        <span slot="left" style="display:inline-block; padding-right: 20px; margin-left: -3px;">请选择单件数量</span>
        <select slot="right" style="text-align: right;" @change="handleNumberSelect">
          <option v-for="(item, index) in weightOptions" :key="index" :value="item">{{ item }}</option>
        </select>
      </yd-cell-item>
    </yd-cell-group>
    <yd-cityselect v-model="showPicker" :callback="handlePick" :items="district"></yd-cityselect>
    <div class="jh-methods">
      <div class="btn-join" :class="[ canJoin ? '' : 'btn-join--disable' ]" @click="handleJoin">报名集货</div>
    </div>
    <pop-box 
      :visible="showPop" 
      :status="status" 
      @update="showPop = false" 
      @btn-click="handlePopClick">
    </pop-box>
  </div>
</template>

<script>
import PopBox from '@/components/PopBox'
import District from 'ydui-district/dist/jd_province_city_area_id'

export default {
  components: { PopBox },
  data () {
    return {
      showPicker: false,
      showPop: false,
      status: false,
      weightOptions: ['0-1kg', '1-2kg', '2-5kg', '5-10kg', '10-20kg', '20kg以上'],
      formData: {
        address: '',
        name: '',
        phone: '',
        weight: '0-1kg',
        number: '',
        detail: ''
      }
    }
  },
  beforeRouteEnter (from, to, next) {
    next(vm => {
      vm.$store.dispatch('SET_TITLE', { title: '填写寄件信息' })
    })
  },
  beforeMount () {
    // this.$store.dispatch('FETCH_ADDRESS')
    //   .catch(err => {
    //     console.log(err)
    //   })
  },
  computed: {
    district () {
      return District || []
    },
    canJoin () {
      for (let name in this.formData) {
        if (!this.formData[name]) {
          return false
        }
      }
      return true
    }
  },
  methods: {
    handlePick (ret) {
      this.formData['address'] = ret.itemName1 + ret.itemName2 + ret.itemName3
    },
    handleJoin () {
      if (!this.canJoin) return false
      this.$store.dispatch('SET_JOIN_DATA', this.formData)
        .then(() => {
          this.showPop = true
        })
        .catch(err => {
          console.log(err)
          this.$dialog.confirm({
            title: '出错了',
            mes: err.toString(),
            opts: () => {
              this.$router.push({ path: '/' })
            }
          })
        })
    },
    handlePopClick () {
      this.$router.push({ path: '/share' })
    },
    handleClose () {
      this.showPop = false
    },
    handleNumberSelect (e) {
      this.formData['weight'] = e.target.value
    }
  }
}
</script>
