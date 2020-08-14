const app = new Vue({
  el: '#app',
  vuetify: new Vuetify(),
  data: {
    form: {
      name: "",
      address: "",
      tel: "",
    },
    headers: [
      {text: "名前", value: "name"},
      {text: "住所", value: "address"},
      {text: "電話番号", value: "tel"},
      {text: "アクション", value: "actions"},
    ],
    items: [],
    item: {},
    dialog: false,
  },
  methods: {
    openDialog(item) {
      this.item.name  = item.name;
      this.item.address  = item.address;
      this.item.tel  = item.tel;
      this.item.id = item.id;
      this.dialog = true;
    },
    registerAddress() {
      axios.post("/address", this.form).then(() => {
        alert("登録しました");
        this.getAllAddress();
      }).catch(error => {
        console.log(error);
        alert(error.response.data.message);
      })
    },
    getAllAddress() {
      axios.get("/address").then(response => {
        this.items = response.data;
      }).catch(error => {
        console.log(error);
        alert(error.response.data.message);
      })
    },
    updateAddress() {
      axios.put(`/address/${this.item.id}`, this.item).then(() => {
        alert("更新しました");
        this.getAllAddress();
      }).catch(error => {
        console.log(error);
        alert(error.response.data.message);
      })
    },
    deleteAddress(id) {
      axios.delete(`/address/${id}`).then(() => {
        alert("削除しました");
        this.getAllAddress();
      }).catch(error => {
        console.log(error);
        alert(error.response.data.message);
      })
    },
  },
  mounted(){
    this.getAllAddress();
  }
})

