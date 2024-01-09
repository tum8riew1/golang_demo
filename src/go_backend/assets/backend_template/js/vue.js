Vue.component('groups-datatable-components', {
    mounted() {
        this.fetchData;
    },
    props: {
        app_url: '',
        api_url: '',
        api_path_data: '',
        api_path_delete: '',
        url_add_data: '',
        url_edit: '',
        token: '',
    },
    data() {
        return {
            search: '',
            headers: [{
                    text: 'Name',
                    align: 'start',
                    sortable: false,
                    value: 'name',
                },
                { text: 'Crated At', value: 'created_at' },
                { text: 'Status', value: 'status' },
                { text: 'Actions', value: 'actions', sortable: false },
            ],
            data_list: [],
        }
    },
    computed: {
        fetchData: async function(e) {
            //console.log("FetchData", this.api_url, this.api_path_data, this.token)
            const params_user_session = {};
            response = await axios.post(this.api_url + this.api_path_data, params_user_session, {
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': 'Bearer ' + this.token
                },
            }).catch(function(error) {
                console.log('Error! Could not reach the API. ' + error);
            });
            //console.log(response.status, response.data);
            if (response.status == 200) {
                this.data_list = response.data.data
            }
        },
    },
    methods: {
        editItem(item) {
            console.log(item.id);
            location.href = this.app_url + this.url_edit + '/' + item.id
        },
        deleteItem: async function(item) {
            console.log(item);
            response = await axios.get(this.api_url + this.api_path_delete + '/' + item.id, {
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': 'Bearer ' + this.token
                },
            }).catch(function(error) {
                console.log('Error! Could not reach the API. ' + error);
            });
            console.log(response.status, response.data);
            if (response.status == 200) {
                //console.log("call fetch");
                this.fetchDataUpdate()
            }
        },
        addItem() {
            //console.log("add Item", this.api_url, this.url_add_data)
            location.href = this.app_url + this.url_add_data
        },
        fetchDataUpdate: async function(e) {
            //console.log("fetchDataUpdate", this.api_url, this.api_path_data, this.token)
            const params_user_session = {};
            response = await axios.post(this.api_url + this.api_path_data, params_user_session, {
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': 'Bearer ' + this.token
                },
            }).catch(function(error) {
                console.log('Error! Could not reach the API. ' + error);
            });
            //console.log(response.status, response.data);
            if (response.status == 200) {
                this.data_list = response.data.data
            }
        }
    },
    template: `
    <v-card>
      <v-card-title>
        
        <v-text-field
          v-model="search"
          append-icon="mdi-magnify"
          label="Search"
          single-line
          hide-details
        ></v-text-field>
        <v-spacer></v-spacer>

        <v-btn
            color="primary"
            dark
            class="mb-2"
            @click="addItem()"
            >
            New Item
        </v-btn>

      </v-card-title>
      <v-data-table
        :headers="headers"
        :items="data_list"
        :search="search"
      >
        <template v-slot:item.actions="{ item }">
            <v-icon
                small
                class="mr-2"
                @click="editItem(item)"
            >
                mdi-pencil
            </v-icon>
            <v-icon
                small
                @click="deleteItem(item)"
            >
                mdi-delete
            </v-icon>
        </template>
      </v-data-table>
    </v-card>`
})


Vue.component('groups-add-components', {
    mounted() {
        //console.log(this.app_url, this.api_url, this.token);
    },
    props: {
        app_url: '',
        api_url: '',
        url_store: '',
        url_list: '',
        token: '',
        label_name: '',
        label_status: '',
    },
    data() {
        return {
            name: '',
            status: '1',
        }
    },
    computed: {

    },
    methods: {
        validateBeforeSubmit() {
            console.log("Submit");
            //this.$refs.observer.validate()
            this.$validator.validateAll().then((result) => {
                if (result) {
                    // eslint-disable-next-line
                    console.log('Form Submitted!');

                    const params = {
                        name: this.name,
                        status: this.status
                    };
                    axios.post(this.api_url + this.url_store, params, {
                        headers: {
                            'Content-Type': 'application/json',
                            'Authorization': 'Bearer ' + this.token
                        },
                    }).then(response => {

                        console.log(response);
                        //loader.hide();

                        if (response.status == 200) {
                            //   loader.hide();
                            //   Vue.swal({
                            //     position: 'center',
                            //     type: 'success',
                            //     title:this.message_save_data,
                            //     showConfirmButton: false,
                            //     timer: 5000
                            //   });
                            location.href = this.app_url + this.url_list;
                        }

                    }).catch(function(error) {
                        console.log('Error! Could not reach the API. ' + error);
                    });
                    return;
                }
                console.log('Correct them errors!');
            });
        },
        cancel() {
            location.href = this.app_url + this.url_list;
        },
    },
    //form-control
    template: `
    <form @submit.prevent="validateBeforeSubmit">
        <div class="form-row">
            <div class="form-group col-md-6">
                <label for="inputName">{{ label_name }}</label>
                    <input name="name" v-model="name" v-validate="'required'" :class="{'form-control': true, 'is-danger': errors.has('name') }" type="text" placeholder="Name">
                    <i v-show="errors.has('name')" class="fa fa-warning" style="color:red"></i>
                    <span v-show="errors.has('name')" class="help is-danger">{{ errors.first('name') }}</span>
            </div>
        </div>
        <div class="form-row">
            <div class="form-group col-md-6">
                <label for="inputName">{{ label_status }}</label>
                <v-radio-group v-model="status">
                    <v-radio
                        label="Publish"
                        color="primary"
                        value="1"
                    ></v-radio>
                    <v-radio
                    label="Unpublish"
                    color="error"
                    value="0"
                    ></v-radio>
              </v-radio-group>
            </div>
        </div>
        <v-btn
        depressed
        color="primary"
        type="submit"
        >
            Submit
        </v-btn>
        <v-btn depressed @click="cancel(item)">
            Cancel
        </v-btn>
    </form>`
})


Vue.component('groups-edit-components', {
    mounted() {
        //console.log(this.app_url, this.api_url, this.token, this.id);
        this.fetchData
    },
    props: {
        app_url: '',
        api_url: '',
        url_update: '',
        url_data: '',
        url_list: '',
        token: '',
        id: '',
        label_name: '',
        label_status: '',
    },
    data() {
        return {
            name: '',
            status: '1',
        }
    },
    computed: {
        fetchData: async function(e) {
            console.log("FetchData", this.api_url, this.api_path_data, this.token)

            response = await axios.get(this.api_url + this.url_data + '/' + this.id, {
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': 'Bearer ' + this.token
                },
            }).catch(function(error) {
                console.log('Error! Could not reach the API. ' + error);
            });
            //console.log(response.status, response.data);
            if (response.status == 200) {
                this.name = response.data.data.name
                this.status = String(response.data.data.status)
            }
        },
    },
    methods: {
        validateBeforeSubmit() {
            console.log("Submit");
            //this.$refs.observer.validate()
            this.$validator.validateAll().then((result) => {
                if (result) {
                    // eslint-disable-next-line
                    console.log('Form Submitted!');

                    const params = {
                        id: Number(this.id),
                        name: this.name,
                        status: this.status
                    };
                    console.log(params)
                    axios.post(this.api_url + this.url_update, params, {
                        headers: {
                            'Content-Type': 'application/json',
                            'Authorization': 'Bearer ' + this.token
                        },
                    }).then(response => {

                        console.log(response);
                        //loader.hide();

                        if (response.status == 200) {
                            //   loader.hide();
                            //   Vue.swal({
                            //     position: 'center',
                            //     type: 'success',
                            //     title:this.message_save_data,
                            //     showConfirmButton: false,
                            //     timer: 5000
                            //   });
                            location.href = this.app_url + this.url_list;
                        }

                    }).catch(function(error) {
                        console.log('Error! Could not reach the API. ' + error);
                    });
                    return;
                }
                console.log('Correct them errors!');
            });
        },
        cancel() {
            location.href = this.app_url + this.url_list;
        },
    },
    //form-control
    template: `
    <form @submit.prevent="validateBeforeSubmit">
        <div class="form-row">
            <div class="form-group col-md-6">
                <label for="inputName">{{ label_name }}</label>
                    <input name="name" v-model="name" v-validate="'required'" :class="{'form-control': true, 'is-danger': errors.has('name') }" type="text" placeholder="Name">
                    <i v-show="errors.has('name')" class="fa fa-warning" style="color:red"></i>
                    <span v-show="errors.has('name')" class="help is-danger">{{ errors.first('name') }}</span>
            </div>
        </div>
        <div class="form-row">
            <div class="form-group col-md-6">
                <label for="inputName">{{ label_status }}</label>
                <v-radio-group v-model="status">
                    <v-radio
                        label="Publish"
                        color="primary"
                        value="1"
                    ></v-radio>
                    <v-radio
                    label="Unpublish"
                    color="error"
                    value="0"
                    ></v-radio>
              </v-radio-group>
            </div>
        </div>
        <v-btn
        depressed
        color="primary"
        type="submit"
        >
            Submit
        </v-btn>
        <v-btn depressed @click="cancel(item)">
            Cancel
        </v-btn>
    </form>`
})


Vue.component('users-datatable-components', {
    mounted() {
        this.fetchData;
    },
    props: {
        app_url: '',
        api_url: '',
        api_path_data: '',
        api_path_delete: '',
        url_add_data: '',
        url_edit: '',
        token: '',
    },
    data() {
        return {
            search: '',
            headers: [{
                    text: 'Name',
                    align: 'start',
                    sortable: false,
                    value: 'name',
                },
                { text: 'Group', value: 'role_name' },
                { text: 'Crated At', value: 'created_at' },
                { text: 'Status', value: 'status' },
                { text: 'Actions', value: 'actions', sortable: false },
            ],
            data_list: [],
        }
    },
    computed: {
        fetchData: async function(e) {
            //console.log("FetchData", this.api_url, this.api_path_data, this.token)
            const params_user_session = {};
            response = await axios.post(this.api_url + this.api_path_data, params_user_session, {
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': 'Bearer ' + this.token
                },
            }).catch(function(error) {
                console.log('Error! Could not reach the API. ' + error);
            });
            //console.log(response.status, response.data);
            if (response.status == 200) {
                this.data_list = response.data.data
            }
        },
    },
    methods: {
        editItem(item) {
            console.log(item.id);
            location.href = this.app_url + this.url_edit + '/' + item.id
        },
        deleteItem: async function(item) {
            console.log(item);
            response = await axios.get(this.api_url + this.api_path_delete + '/' + item.id, {
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': 'Bearer ' + this.token
                },
            }).catch(function(error) {
                console.log('Error! Could not reach the API. ' + error);
            });
            console.log(response.status, response.data);
            if (response.status == 200) {
                //console.log("call fetch");
                this.fetchDataUpdate()
            }
        },
        addItem() {
            //console.log("add Item", this.api_url, this.url_add_data)
            location.href = this.app_url + this.url_add_data
        },
        fetchDataUpdate: async function(e) {
            //console.log("fetchDataUpdate", this.api_url, this.api_path_data, this.token)
            const params_user_session = {};
            response = await axios.post(this.api_url + this.api_path_data, params_user_session, {
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': 'Bearer ' + this.token
                },
            }).catch(function(error) {
                console.log('Error! Could not reach the API. ' + error);
            });
            //console.log(response.status, response.data);
            if (response.status == 200) {
                this.data_list = response.data.data
            }
        }
    },
    template: `
    <v-card>
      <v-card-title>
        
        <v-text-field
          v-model="search"
          append-icon="mdi-magnify"
          label="Search"
          single-line
          hide-details
        ></v-text-field>
        <v-spacer></v-spacer>

        <v-btn
            color="primary"
            dark
            class="mb-2"
            @click="addItem()"
            >
            New Item
        </v-btn>

      </v-card-title>
      <v-data-table
        :headers="headers"
        :items="data_list"
        :search="search"
      >
        <template v-slot:item.actions="{ item }">
            <v-icon
                small
                class="mr-2"
                @click="editItem(item)"
            >
                mdi-pencil
            </v-icon>
            <v-icon
                small
                @click="deleteItem(item)"
            >
                mdi-delete
            </v-icon>
        </template>
      </v-data-table>
    </v-card>`
})

Vue.component('users-add-components', {
    mounted() {
        //console.log(this.app_url, this.api_url, this.token);
        this.fetchGroupsData;
    },
    props: {
        app_url: '',
        api_url: '',
        url_store: '',
        url_groups: '',
        url_list: '',
        token: '',
        label_name: '',
        label_username: '',
        label_email: '',
        label_group: '',
        label_password: '',
        label_status: '',
        user_id: '',
    },
    data() {
        return {
            name: '',
            username: '',
            email: '',
            group: '',
            password: '',
            status: '1',
            group_items: [],
        }
    },
    computed: {
        fetchGroupsData: async function(e) {
            //console.log("FetchData", this.api_url, this.api_path_data, this.token)
            const params_user_session = {};
            response = await axios.post(this.api_url + this.url_groups, params_user_session, {
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': 'Bearer ' + this.token
                },
            }).catch(function(error) {
                console.log('Error! Could not reach the API. ' + error);
            });
            console.log(response.status, response.data);
            if (response.status == 200) {
                this.group_items = response.data.data
            }
        },
    },
    methods: {
        validateBeforeSubmit() {
            console.log("Submit");
            //this.$refs.observer.validate()
            this.$validator.validateAll().then((result) => {
                if (result) {
                    // eslint-disable-next-line
                    console.log('Form Submitted!');

                    const params = {
                        name: this.name,
                        username: this.username,
                        password: this.password,
                        email: this.email,
                        role_id: this.group,
                        created_by: parseInt(this.user_id),
                        status: this.status
                    };
                    console.log(params);
                    axios.post(this.api_url + this.url_store, params, {
                        headers: {
                            'Content-Type': 'application/json',
                            'Authorization': 'Bearer ' + this.token
                        },
                    }).then(response => {

                        console.log(response);
                        //loader.hide();

                        if (response.status == 200) {
                            //   loader.hide();
                            //   Vue.swal({
                            //     position: 'center',
                            //     type: 'success',
                            //     title:this.message_save_data,
                            //     showConfirmButton: false,
                            //     timer: 5000
                            //   });
                            location.href = this.app_url + this.url_list;
                        }

                    }).catch(function(error) {
                        console.log('Error! Could not reach the API. ' + error);
                    });
                    return;
                }
                console.log('Correct them errors!');
            });
        },
        cancel() {
            location.href = this.app_url + this.url_list;
        },
    },
    //form-control
    template: `
    <form @submit.prevent="validateBeforeSubmit">
        <div class="form-row">
            <div class="form-group col-md-6">
                <label for="inputUserName">{{ label_username }}</label>
                    <input name="username" v-model="username" v-validate="'required'" :class="{'form-control': true, 'is-danger': errors.has('username') }" type="text" v-bind:placeholder="label_username" >
                    <i v-show="errors.has('username')" class="fa fa-warning" style="color:red"></i>
                    <span v-show="errors.has('username')" class="help is-danger">{{ errors.first('username') }}</span>
            </div>
        </div>
        <div class="form-row">
            <div class="form-group col-md-6">
                <label for="inputName">{{ label_name }}</label>
                    <input name="name" v-model="name" v-validate="'required'" :class="{'form-control': true, 'is-danger': errors.has('name') }" type="text" v-bind:placeholder="label_name">
                    <i v-show="errors.has('name')" class="fa fa-warning" style="color:red"></i>
                    <span v-show="errors.has('name')" class="help is-danger">{{ errors.first('name') }}</span>
            </div>
        </div>
        <div class="form-row">
            <div class="form-group col-md-6">
                <label for="inputEmail">{{ label_email }}</label>
                    <input name="email" v-model="email" v-validate="'required|email'" :class="{'form-control': true, 'is-danger': errors.has('name') }" type="text"  v-bind:placeholder="label_email">
                    <i v-show="errors.has('email')" class="fa fa-warning" style="color:red"></i>
                    <span v-show="errors.has('email')" class="help is-danger">{{ errors.first('email') }}</span>
            </div>
        </div>
        <div class="form-row">
            <div class="form-group col-md-6">
                    <v-select
                        :items="group_items"
                        item-text="name"
                        item-value="id"
                        name="group" v-model="group" v-validate="'required'"
                        v-bind:label="label_group"
                    ></v-select>
                    <i v-show="errors.has('group')" class="fa fa-warning" style="color:red"></i>
                    <span v-show="errors.has('group')" class="help is-danger">{{ errors.first('group') }}</span>
            </div>
        </div>
        <div class="form-row">
            <div class="form-group col-md-6">
                <label for="inputPaswword">{{ label_password }}</label>
                    <input name="password" v-model="password" v-validate="'required|min:4|max:60'" :class="{'form-control': true, 'is-danger': errors.has('name') }" type="password"  v-bind:placeholder="label_password">
                    <i v-show="errors.has('password')" class="fa fa-warning" style="color:red"></i>
                    <span v-show="errors.has('password')" class="help is-danger">{{ errors.first('password') }}</span>
            </div>
        </div>        
        <div class="form-row">
            <div class="form-group col-md-6">
                <label for="inputName">{{ label_status }}</label>
                <v-radio-group v-model="status">
                    <v-radio
                        label="Publish"
                        color="primary"
                        value="1"
                    ></v-radio>
                    <v-radio
                    label="Unpublish"
                    color="error"
                    value="0"
                    ></v-radio>
              </v-radio-group>
            </div>
        </div>
        <v-btn
        depressed
        color="primary"
        type="submit"
        >
            Submit
        </v-btn>
        <v-btn depressed @click="cancel(item)">
            Cancel
        </v-btn>
    </form>`
})

Vue.component('users-edit-components', {
    async mounted() {
        //console.log(this.app_url, this.api_url, this.token);
        await this.fetchGroupsData;
        this.fetchData;
    },
    props: {
        app_url: '',
        api_url: '',
        url_update: '',
        url_groups: '',
        url_data: '',
        url_list: '',
        token: '',
        label_name: '',
        label_username: '',
        label_email: '',
        label_group: '',
        label_password: '',
        label_status: '',
        user_id: '',
        id: '',
    },
    data() {
        return {
            name: '',
            username: '',
            email: '',
            group: '',
            password: '',
            status: '1',
            group_items: [],
        }
    },
    computed: {
        fetchGroupsData: async function(e) {
            //console.log("FetchData", this.api_url, this.api_path_data, this.token)
            const params_user_session = {};
            response = await axios.post(this.api_url + this.url_groups, params_user_session, {
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': 'Bearer ' + this.token
                },
            }).catch(function(error) {
                console.log('Error! Could not reach the API. ' + error);
            });
            console.log(response.status, response.data);
            if (response.status == 200) {
                this.group_items = response.data.data
            }
        },
        fetchData: async function(e) {
            console.log("FetchData", this.api_url, this.url_data, this.token)

            response = await axios.get(this.api_url + this.url_data + '/' + this.id, {
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': 'Bearer ' + this.token
                },
            }).catch(function(error) {
                console.log('Error! Could not reach the API. ' + error);
            });
            console.log(response.status, response.data);
            if (response.status == 200) {
                this.username = response.data.data.username
                this.name = response.data.data.name
                this.email = response.data.data.email
                this.group = response.data.data.role_id
                this.status = String(response.data.data.status)
            }
        }
    },
    methods: {
        validateBeforeSubmit() {
            console.log("Submit");
            //this.$refs.observer.validate()
            this.$validator.validateAll().then((result) => {
                if (result) {
                    // eslint-disable-next-line
                    console.log('Form Submitted!');

                    const params = {
                        id: parseInt(this.id),
                        name: this.name,
                        password: this.password,
                        role_id: this.group,
                        updated_by: parseInt(this.user_id),
                        status: String(this.status)
                    };
                    console.log(params);
                    axios.post(this.api_url + this.url_update, params, {
                        headers: {
                            'Content-Type': 'application/json',
                            'Authorization': 'Bearer ' + this.token
                        },
                    }).then(response => {

                        console.log(response);
                        //loader.hide();

                        if (response.status == 200) {
                            //   loader.hide();
                            //   Vue.swal({
                            //     position: 'center',
                            //     type: 'success',
                            //     title:this.message_save_data,
                            //     showConfirmButton: false,
                            //     timer: 5000
                            //   });
                            location.href = this.app_url + this.url_list;
                        }

                    }).catch(function(error) {
                        console.log('Error! Could not reach the API. ' + error);
                    });
                    return;
                }
                console.log('Correct them errors!');
            });
        },
        cancel() {
            location.href = this.app_url + this.url_list;
        },
    },
    //form-control
    template: `
    <form @submit.prevent="validateBeforeSubmit">
        <div class="form-row">
            <div class="form-group col-md-6">
                <label for="inputUserName">{{ label_username }}</label>
                    <input name="username" v-model="username" v-validate="'required'" :class="{'form-control': true, 'is-danger': errors.has('username') }" type="text" v-bind:placeholder="label_username" readonly>
                    <i v-show="errors.has('username')" class="fa fa-warning" style="color:red"></i>
                    <span v-show="errors.has('username')" class="help is-danger">{{ errors.first('username') }}</span>
            </div>
        </div>
        <div class="form-row">
            <div class="form-group col-md-6">
                <label for="inputName">{{ label_name }}</label>
                    <input name="name" v-model="name" v-validate="'required'" :class="{'form-control': true, 'is-danger': errors.has('name') }" type="text" v-bind:placeholder="label_name">
                    <i v-show="errors.has('name')" class="fa fa-warning" style="color:red"></i>
                    <span v-show="errors.has('name')" class="help is-danger">{{ errors.first('name') }}</span>
            </div>
        </div>
        <div class="form-row">
            <div class="form-group col-md-6">
                <label for="inputEmail">{{ label_email }}</label>
                    <input name="email" v-model="email" v-validate="'required|email'" :class="{'form-control': true, 'is-danger': errors.has('name') }" type="text"  v-bind:placeholder="label_email" readonly>
                    <i v-show="errors.has('email')" class="fa fa-warning" style="color:red"></i>
                    <span v-show="errors.has('email')" class="help is-danger">{{ errors.first('email') }}</span>
            </div>
        </div>
        <div class="form-row">
            <div class="form-group col-md-6">
                    <v-select
                        :items="group_items"
                        item-text="name"
                        item-value="id"
                        name="group" v-model="group" v-validate="'required'"
                        v-bind:label="label_group"
                    ></v-select>
                    <i v-show="errors.has('group')" class="fa fa-warning" style="color:red"></i>
                    <span v-show="errors.has('group')" class="help is-danger">{{ errors.first('group') }}</span>
            </div>
        </div>
        <div class="form-row">
            <div class="form-group col-md-6">
                <label for="inputPaswword">{{ label_password }}</label>
                    <input name="password" v-model="password" v-validate="'min:4|max:60'" :class="{'form-control': true, 'is-danger': errors.has('name') }" type="password"  v-bind:placeholder="label_password">
                    <i v-show="errors.has('password')" class="fa fa-warning" style="color:red"></i>
                    <span v-show="errors.has('password')" class="help is-danger">{{ errors.first('password') }}</span>
            </div>
        </div>        
        <div class="form-row">
            <div class="form-group col-md-6">
                <label for="inputName">{{ label_status }}</label>
                <v-radio-group v-model="status">
                    <v-radio
                        label="Publish"
                        color="primary"
                        value="1"
                    ></v-radio>
                    <v-radio
                    label="Unpublish"
                    color="error"
                    value="0"
                    ></v-radio>
              </v-radio-group>
            </div>
        </div>
        <v-btn
        depressed
        color="primary"
        type="submit"
        >
            Submit
        </v-btn>
        <v-btn depressed @click="cancel(item)">
            Cancel
        </v-btn>
    </form>`
})



Vue.component('categorys-datatable-components', {
    mounted() {
        this.fetchData;
    },
    props: {
        app_url: '',
        api_url: '',
        api_path_data: '',
        api_path_delete: '',
        url_add_data: '',
        url_edit: '',
        token: '',
    },
    data() {
        return {
            search: '',
            headers: [{
                    text: 'Title',
                    align: 'start',
                    sortable: false,
                    value: 'title',
                },
                { text: 'Crated At', value: 'created_at' },
                { text: 'Status', value: 'status' },
                { text: 'Actions', value: 'actions', sortable: false },
            ],
            data_list: [],
        }
    },
    computed: {
        fetchData: async function(e) {
            //console.log("FetchData", this.api_url, this.api_path_data, this.token)
            const params_user_session = {};
            response = await axios.post(this.api_url + this.api_path_data, params_user_session, {
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': 'Bearer ' + this.token
                },
            }).catch(function(error) {
                console.log('Error! Could not reach the API. ' + error);
            });
            //console.log(response.status, response.data);
            if (response.status == 200) {
                this.data_list = response.data.data
            }
        },
    },
    methods: {
        editItem(item) {
            console.log(item.id);
            location.href = this.app_url + this.url_edit + '/' + item.id
        },
        deleteItem: async function(item) {
            console.log(item);
            response = await axios.get(this.api_url + this.api_path_delete + '/' + item.id, {
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': 'Bearer ' + this.token
                },
            }).catch(function(error) {
                console.log('Error! Could not reach the API. ' + error);
            });
            console.log(response.status, response.data);
            if (response.status == 200) {
                //console.log("call fetch");
                this.fetchDataUpdate()
            }
        },
        addItem() {
            //console.log("add Item", this.api_url, this.url_add_data)
            location.href = this.app_url + this.url_add_data
        },
        fetchDataUpdate: async function(e) {
            //console.log("fetchDataUpdate", this.api_url, this.api_path_data, this.token)
            const params_user_session = {};
            response = await axios.post(this.api_url + this.api_path_data, params_user_session, {
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': 'Bearer ' + this.token
                },
            }).catch(function(error) {
                console.log('Error! Could not reach the API. ' + error);
            });
            //console.log(response.status, response.data);
            if (response.status == 200) {
                this.data_list = response.data.data
            }
        }
    },
    template: `
    <v-card>
      <v-card-title>
        
        <v-text-field
          v-model="search"
          append-icon="mdi-magnify"
          label="Search"
          single-line
          hide-details
        ></v-text-field>
        <v-spacer></v-spacer>

        <v-btn
            color="primary"
            dark
            class="mb-2"
            @click="addItem()"
            >
            New Item
        </v-btn>

      </v-card-title>
      <v-data-table
        :headers="headers"
        :items="data_list"
        :search="search"
      >
        <template v-slot:item.actions="{ item }">
            <v-icon
                small
                class="mr-2"
                @click="editItem(item)"
            >
                mdi-pencil
            </v-icon>
            <v-icon
                small
                @click="deleteItem(item)"
            >
                mdi-delete
            </v-icon>
        </template>
      </v-data-table>
    </v-card>`
})

Vue.component('categorys-add-components', {
    mounted() {
        //console.log(this.app_url, this.api_url, this.token);
    },
    props: {
        app_url: '',
        api_url: '',
        url_store: '',
        url_list: '',
        token: '',
        label_name: '',
        label_description: '',
        label_short_description: '',
        label_status: '',
        user_id: ''
    },
    data() {
        return {
            name: '',
            description: '',
            short_description: '',
            status: '1',
        }
    },
    computed: {

    },
    methods: {
        validateBeforeSubmit() {
            console.log("Submit");
            //this.$refs.observer.validate()
            this.$validator.validateAll().then((result) => {
                if (result) {
                    // eslint-disable-next-line
                    console.log('Form Submitted!');

                    const params = {
                        title: this.name,
                        description: this.description,
                        short_description: this.short_description,
                        created_by: parseInt(this.user_id),
                        status: this.status
                    };
                    axios.post(this.api_url + this.url_store, params, {
                        headers: {
                            'Content-Type': 'application/json',
                            'Authorization': 'Bearer ' + this.token
                        },
                    }).then(response => {

                        console.log(response);
                        //loader.hide();

                        if (response.status == 200) {
                            //   loader.hide();
                            //   Vue.swal({
                            //     position: 'center',
                            //     type: 'success',
                            //     title:this.message_save_data,
                            //     showConfirmButton: false,
                            //     timer: 5000
                            //   });
                            location.href = this.app_url + this.url_list;
                        }

                    }).catch(function(error) {
                        console.log('Error! Could not reach the API. ' + error);
                    });
                    return;
                }
                console.log('Correct them errors!');
            });
        },
        cancel() {
            location.href = this.app_url + this.url_list;
        },
    },
    //form-control
    template: `
    <form @submit.prevent="validateBeforeSubmit">
        <div class="form-row">
            <div class="form-group col-md-6">
                <label for="inputName">{{ label_name }}</label>
                <input name="name" v-model="name" v-validate="'required'" :class="{'form-control': true, 'is-danger': errors.has('name') }" type="text" v-bind:placeholder="label_name">
                <i v-show="errors.has('name')" class="fa fa-warning" style="color:red"></i>
                <span v-show="errors.has('name')" class="help is-danger">{{ errors.first('name') }}</span>
            </div>
        </div>
        <div class="form-row">
            <div class="form-group col-md-6">
                <label for="inputShort_descriptio">{{ label_short_description }}</label>
                <textarea id="short_description" v-model="short_description" v-bind:placeholder="label_short_description" class="form-control" cols="50" rows="10"></textarea>
                <i v-show="errors.has('short_description')" class="fa fa-warning" style="color:red"></i>
                <span v-show="errors.has('short_description')" class="help is-danger">{{ errors.first('short_description') }}</span>
            </div>
        </div>
        <div class="form-row">
            <div class="form-group col-md-6">
                <label for="inputDescription">{{ label_description }}</label>
                <textarea id="description" v-model="description" v-bind:placeholder="label_description" class="form-control" cols="50" rows="10"></textarea>
                <i v-show="errors.has('description')" class="fa fa-warning" style="color:red"></i>
                <span v-show="errors.has('description')" class="help is-danger">{{ errors.first('description') }}</span>
            </div>
        </div>
        <div class="form-row">
            <div class="form-group col-md-6">
                <label for="inputStatus">{{ label_status }}</label>
                <v-radio-group v-model="status">
                    <v-radio
                        label="Publish"
                        color="primary"
                        value="1"
                    ></v-radio>
                    <v-radio
                    label="Unpublish"
                    color="error"
                    value="0"
                    ></v-radio>
              </v-radio-group>
            </div>
        </div>
        <v-btn
        depressed
        color="primary"
        type="submit"
        >
            Submit
        </v-btn>
        <v-btn depressed @click="cancel(item)">
            Cancel
        </v-btn>
    </form>`
})

Vue.component('categorys-edit-components', {
    mounted() {
        //console.log(this.app_url, this.api_url, this.token);
        this.fetchData
    },
    props: {
        app_url: '',
        api_url: '',
        url_update: '',
        url_data: '',
        url_list: '',
        token: '',
        label_name: '',
        label_description: '',
        label_short_description: '',
        label_status: '',
        user_id: '',
        id: '',
    },
    data() {
        return {
            name: '',
            description: '',
            short_description: '',
            status: '1',
        }
    },
    computed: {
        fetchData: async function(e) {
            console.log("FetchData", this.api_url, this.url_data, this.token)

            response = await axios.get(this.api_url + this.url_data + '/' + this.id, {
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': 'Bearer ' + this.token
                },
            }).catch(function(error) {
                console.log('Error! Could not reach the API. ' + error);
            });
            console.log(response.status, response.data);
            if (response.status == 200) {
                this.name = response.data.data.title
                this.description = response.data.data.description
                this.short_description = response.data.data.short_description
                this.status = String(response.data.data.status)
            }
        }
    },
    methods: {
        validateBeforeSubmit() {
            console.log("Submit");
            //this.$refs.observer.validate()
            this.$validator.validateAll().then((result) => {
                if (result) {
                    // eslint-disable-next-line
                    console.log('Form Submitted!');

                    const params = {
                        id: parseInt(this.id),
                        title: this.name,
                        description: this.description,
                        short_description: this.short_description,
                        updated_by: parseInt(this.user_id),
                        status: this.status
                    };
                    console.log(params);
                    axios.post(this.api_url + this.url_update, params, {
                        headers: {
                            'Content-Type': 'application/json',
                            'Authorization': 'Bearer ' + this.token
                        },
                    }).then(response => {

                        console.log(response);
                        //loader.hide();

                        if (response.status == 200) {
                            //   loader.hide();
                            //   Vue.swal({
                            //     position: 'center',
                            //     type: 'success',
                            //     title:this.message_save_data,
                            //     showConfirmButton: false,
                            //     timer: 5000
                            //   });
                            location.href = this.app_url + this.url_list;
                        }

                    }).catch(function(error) {
                        console.log('Error! Could not reach the API. ' + error);
                    });
                    return;
                }
                console.log('Correct them errors!');
            });
        },
        cancel() {
            location.href = this.app_url + this.url_list;
        },
    },
    //form-control
    template: `
    <form @submit.prevent="validateBeforeSubmit">
        <div class="form-row">
            <div class="form-group col-md-6">
                <label for="inputName">{{ label_name }}</label>
                <input name="name" v-model="name" v-validate="'required'" :class="{'form-control': true, 'is-danger': errors.has('name') }" type="text" v-bind:placeholder="label_name">
                <i v-show="errors.has('name')" class="fa fa-warning" style="color:red"></i>
                <span v-show="errors.has('name')" class="help is-danger">{{ errors.first('name') }}</span>
            </div>
        </div>
        <div class="form-row">
            <div class="form-group col-md-6">
                <label for="inputShort_descriptio">{{ label_short_description }}</label>
                <textarea id="short_description" v-model="short_description" v-bind:placeholder="label_short_description" class="form-control" cols="50" rows="10"></textarea>
                <i v-show="errors.has('short_description')" class="fa fa-warning" style="color:red"></i>
                <span v-show="errors.has('short_description')" class="help is-danger">{{ errors.first('short_description') }}</span>
            </div>
        </div>
        <div class="form-row">
            <div class="form-group col-md-6">
                <label for="inputDescription">{{ label_description }}</label>
                <textarea id="description" v-model="description" v-bind:placeholder="label_description" class="form-control" cols="50" rows="10"></textarea>
                <i v-show="errors.has('description')" class="fa fa-warning" style="color:red"></i>
                <span v-show="errors.has('description')" class="help is-danger">{{ errors.first('description') }}</span>
            </div>
        </div>
        <div class="form-row">
            <div class="form-group col-md-6">
                <label for="inputStatus">{{ label_status }}</label>
                <v-radio-group v-model="status">
                    <v-radio
                        label="Publish"
                        color="primary"
                        value="1"
                    ></v-radio>
                    <v-radio
                    label="Unpublish"
                    color="error"
                    value="0"
                    ></v-radio>
              </v-radio-group>
            </div>
        </div>
        <v-btn
        depressed
        color="primary"
        type="submit"
        >
            Submit
        </v-btn>
        <v-btn depressed @click="cancel(item)">
            Cancel
        </v-btn>
    </form>`
})

Vue.component('articles-datatable-components', {
    mounted() {
        this.fetchData;
    },
    props: {
        app_url: '',
        api_url: '',
        api_path_data: '',
        api_path_delete: '',
        url_add_data: '',
        url_edit: '',
        token: '',
    },
    data() {
        return {
            search: '',
            headers: [{
                    text: 'Title',
                    align: 'start',
                    sortable: false,
                    value: 'title',
                },
                { text: 'Category', value: 'category_name' },
                { text: 'Crated At', value: 'created_at' },
                { text: 'Status', value: 'status' },
                { text: 'Actions', value: 'actions', sortable: false },
            ],
            data_list: [],
        }
    },
    computed: {
        fetchData: async function(e) {
            //console.log("FetchData", this.api_url, this.api_path_data, this.token)
            const params_user_session = {};
            response = await axios.post(this.api_url + this.api_path_data, params_user_session, {
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': 'Bearer ' + this.token
                },
            }).catch(function(error) {
                console.log('Error! Could not reach the API. ' + error);
            });
            console.log(response.status, response.data);
            if (response.status == 200) {
                this.data_list = response.data.data
            }
        },
    },
    methods: {
        editItem(item) {
            console.log(item.id);
            location.href = this.app_url + this.url_edit + '/' + item.id
        },
        deleteItem: async function(item) {
            console.log(item);
            response = await axios.get(this.api_url + this.api_path_delete + '/' + item.id, {
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': 'Bearer ' + this.token
                },
            }).catch(function(error) {
                console.log('Error! Could not reach the API. ' + error);
            });
            console.log(response.status, response.data);
            if (response.status == 200) {
                //console.log("call fetch");
                this.fetchDataUpdate()
            }
        },
        addItem() {
            //console.log("add Item", this.api_url, this.url_add_data)
            location.href = this.app_url + this.url_add_data
        },
        fetchDataUpdate: async function(e) {
            //console.log("fetchDataUpdate", this.api_url, this.api_path_data, this.token)
            const params_user_session = {};
            response = await axios.post(this.api_url + this.api_path_data, params_user_session, {
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': 'Bearer ' + this.token
                },
            }).catch(function(error) {
                console.log('Error! Could not reach the API. ' + error);
            });
            //console.log(response.status, response.data);
            if (response.status == 200) {
                this.data_list = response.data.data
            }
        }
    },
    template: `
    <v-card>
      <v-card-title>
        
        <v-text-field
          v-model="search"
          append-icon="mdi-magnify"
          label="Search"
          single-line
          hide-details
        ></v-text-field>
        <v-spacer></v-spacer>

        <v-btn
            color="primary"
            dark
            class="mb-2"
            @click="addItem()"
            >
            New Item
        </v-btn>

      </v-card-title>
      <v-data-table
        :headers="headers"
        :items="data_list"
        :search="search"
      >
        <template v-slot:item.actions="{ item }">
            <v-icon
                small
                class="mr-2"
                @click="editItem(item)"
            >
                mdi-pencil
            </v-icon>
            <v-icon
                small
                @click="deleteItem(item)"
            >
                mdi-delete
            </v-icon>
        </template>
      </v-data-table>
    </v-card>`
})

Vue.component('articles-add-components', {
    mounted() {
        //console.log(this.app_url, this.api_url, this.token);
        this.fetchCategorysData
    },
    props: {
        app_url: '',
        api_url: '',
        url_store: '',
        url_category: '',
        url_list: '',
        token: '',
        label_name: '',
        label_description: '',
        label_short_description: '',
        label_category: '',
        label_status: '',
        user_id: ''
    },
    data() {
        return {
            name: '',
            description: '',
            short_description: '',
            category_id: '',
            status: '1',
            category_items: [],
        }
    },
    computed: {
        fetchCategorysData: async function(e) {
            //console.log("FetchData", this.api_url, this.api_path_data, this.token)
            const params_user_session = {};
            response = await axios.post(this.api_url + this.url_category, params_user_session, {
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': 'Bearer ' + this.token
                },
            }).catch(function(error) {
                console.log('Error! Could not reach the API. ' + error);
            });
            //console.log(response.status, response.data);
            if (response.status == 200) {
                this.category_items = response.data.data
            }
        },
    },
    methods: {
        validateBeforeSubmit() {
            console.log("Submit");
            //this.$refs.observer.validate()
            this.$validator.validateAll().then((result) => {
                if (result) {
                    // eslint-disable-next-line
                    console.log('Form Submitted!');

                    const params = {
                        title: this.name,
                        description: this.description,
                        short_description: this.short_description,
                        category_id: parseInt(this.category_id),
                        created_by: parseInt(this.user_id),
                        status: this.status
                    };
                    axios.post(this.api_url + this.url_store, params, {
                        headers: {
                            'Content-Type': 'application/json',
                            'Authorization': 'Bearer ' + this.token
                        },
                    }).then(response => {

                        console.log(response);
                        //loader.hide();

                        if (response.status == 200) {
                            //   loader.hide();
                            //   Vue.swal({
                            //     position: 'center',
                            //     type: 'success',
                            //     title:this.message_save_data,
                            //     showConfirmButton: false,
                            //     timer: 5000
                            //   });
                            location.href = this.app_url + this.url_list;
                        }

                    }).catch(function(error) {
                        console.log('Error! Could not reach the API. ' + error);
                    });
                    return;
                }
                console.log('Correct them errors!');
            });
        },
        cancel() {
            location.href = this.app_url + this.url_list;
        },
    },
    //form-control
    template: `
    <form @submit.prevent="validateBeforeSubmit">
        <div class="form-row">
            <div class="form-group col-md-6">
                <label for="inputName">{{ label_name }}</label>
                <input name="name" v-model="name" v-validate="'required'" :class="{'form-control': true, 'is-danger': errors.has('name') }" type="text" v-bind:placeholder="label_name">
                <i v-show="errors.has('name')" class="fa fa-warning" style="color:red"></i>
                <span v-show="errors.has('name')" class="help is-danger">{{ errors.first('name') }}</span>
            </div>
        </div>
        <div class="form-row">
            <div class="form-group col-md-6">
                <label for="inputShort_descriptio">{{ label_short_description }}</label>
                <textarea id="short_description" v-model="short_description" v-bind:placeholder="label_short_description" class="form-control" cols="50" rows="10"></textarea>
                <i v-show="errors.has('short_description')" class="fa fa-warning" style="color:red"></i>
                <span v-show="errors.has('short_description')" class="help is-danger">{{ errors.first('short_description') }}</span>
            </div>
        </div>
        <div class="form-row">
            <div class="form-group col-md-6">
                <label for="inputDescription">{{ label_description }}</label>
                <textarea id="description" v-model="description" v-bind:placeholder="label_description" class="form-control" cols="50" rows="10"></textarea>
                <i v-show="errors.has('description')" class="fa fa-warning" style="color:red"></i>
                <span v-show="errors.has('description')" class="help is-danger">{{ errors.first('description') }}</span>
            </div>
        </div>
        <div class="form-row">
            <div class="form-group col-md-6">
                <v-select
                    :items="category_items"
                    item-text="title"
                    item-value="id"
                    name="category_id" v-model="category_id" v-validate="'required'"
                    v-bind:label="label_category"
                ></v-select>
                <i v-show="errors.has('category_id')" class="fa fa-warning" style="color:red"></i>
                <span v-show="errors.has('category_id')" class="help is-danger">{{ errors.first('category_id') }}</span>
            </div>
        </div>
        <div class="form-row">
            <div class="form-group col-md-6">
                <label for="inputStatus">{{ label_status }}</label>
                <v-radio-group v-model="status">
                    <v-radio
                        label="Publish"
                        color="primary"
                        value="1"
                    ></v-radio>
                    <v-radio
                    label="Unpublish"
                    color="error"
                    value="0"
                    ></v-radio>
              </v-radio-group>
            </div>
        </div>
        <v-btn
        depressed
        color="primary"
        type="submit"
        >
            Submit
        </v-btn>
        <v-btn depressed @click="cancel(item)">
            Cancel
        </v-btn>
    </form>`
})

Vue.component('articles-edit-components', {
    async mounted() {
        //console.log(this.app_url, this.api_url, this.token);
        await this.fetchCategorysData
        this.fetchData
    },
    props: {
        app_url: '',
        api_url: '',
        url_update: '',
        url_data: '',
        url_list: '',
        url_category: '',
        token: '',
        label_name: '',
        label_description: '',
        label_short_description: '',
        label_status: '',
        label_category: '',
        user_id: '',
        id: '',
    },
    data() {
        return {
            name: '',
            description: '',
            short_description: '',
            category_id: '',
            status: '1',
            category_items: [],
        }
    },
    computed: {
        fetchData: async function(e) {
            //console.log("FetchData", this.api_url, this.url_data, this.token)

            response = await axios.get(this.api_url + this.url_data + '/' + this.id, {
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': 'Bearer ' + this.token
                },
            }).catch(function(error) {
                console.log('Error! Could not reach the API. ' + error);
            });
            //console.log(response.status, response.data);
            if (response.status == 200) {
                this.name = response.data.data.title
                this.description = response.data.data.description
                this.short_description = response.data.data.short_description
                this.category_id = response.data.data.category_id
                this.status = String(response.data.data.status)
            }
        },
        fetchCategorysData: async function(e) {
            //console.log("FetchData", this.api_url, this.api_path_data, this.token)
            const params_user_session = {};
            response = await axios.post(this.api_url + this.url_category, params_user_session, {
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': 'Bearer ' + this.token
                },
            }).catch(function(error) {
                console.log('Error! Could not reach the API. ' + error);
            });
            //console.log(response.status, response.data);
            if (response.status == 200) {
                this.category_items = response.data.data
            }
        },
    },
    methods: {
        validateBeforeSubmit() {
            console.log("Submit");
            //this.$refs.observer.validate()
            this.$validator.validateAll().then((result) => {
                if (result) {
                    // eslint-disable-next-line
                    console.log('Form Submitted!');

                    const params = {
                        id: parseInt(this.id),
                        title: this.name,
                        description: this.description,
                        short_description: this.short_description,
                        category_id: parseInt(this.category_id),
                        updated_by: parseInt(this.user_id),
                        status: this.status
                    };
                    console.log(params);
                    axios.post(this.api_url + this.url_update, params, {
                        headers: {
                            'Content-Type': 'application/json',
                            'Authorization': 'Bearer ' + this.token
                        },
                    }).then(response => {

                        console.log(response);
                        //loader.hide();

                        if (response.status == 200) {
                            //   loader.hide();
                            //   Vue.swal({
                            //     position: 'center',
                            //     type: 'success',
                            //     title:this.message_save_data,
                            //     showConfirmButton: false,
                            //     timer: 5000
                            //   });
                            location.href = this.app_url + this.url_list;
                        }

                    }).catch(function(error) {
                        console.log('Error! Could not reach the API. ' + error);
                    });
                    return;
                }
                console.log('Correct them errors!');
            });
        },
        cancel() {
            location.href = this.app_url + this.url_list;
        },
    },
    //form-control
    template: `
    <form @submit.prevent="validateBeforeSubmit">
        <div class="form-row">
            <div class="form-group col-md-6">
                <label for="inputName">{{ label_name }}</label>
                <input name="name" v-model="name" v-validate="'required'" :class="{'form-control': true, 'is-danger': errors.has('name') }" type="text" v-bind:placeholder="label_name">
                <i v-show="errors.has('name')" class="fa fa-warning" style="color:red"></i>
                <span v-show="errors.has('name')" class="help is-danger">{{ errors.first('name') }}</span>
            </div>
        </div>
        <div class="form-row">
            <div class="form-group col-md-6">
                <label for="inputShort_descriptio">{{ label_short_description }}</label>
                <textarea id="short_description" v-model="short_description" v-bind:placeholder="label_short_description" class="form-control" cols="50" rows="10"></textarea>
                <i v-show="errors.has('short_description')" class="fa fa-warning" style="color:red"></i>
                <span v-show="errors.has('short_description')" class="help is-danger">{{ errors.first('short_description') }}</span>
            </div>
        </div>
        <div class="form-row">
            <div class="form-group col-md-6">
                <label for="inputDescription">{{ label_description }}</label>
                <textarea id="description" v-model="description" v-bind:placeholder="label_description" class="form-control" cols="50" rows="10"></textarea>
                <i v-show="errors.has('description')" class="fa fa-warning" style="color:red"></i>
                <span v-show="errors.has('description')" class="help is-danger">{{ errors.first('description') }}</span>
            </div>
        </div>
        <div class="form-row">
            <div class="form-group col-md-6">
                <v-select
                    :items="category_items"
                    item-text="title"
                    item-value="id"
                    name="category_id" v-model="category_id" v-validate="'required'"
                    v-bind:label="label_category"
                ></v-select>
                <i v-show="errors.has('category_id')" class="fa fa-warning" style="color:red"></i>
                <span v-show="errors.has('category_id')" class="help is-danger">{{ errors.first('category_id') }}</span>
            </div>
        </div>        
        <div class="form-row">
            <div class="form-group col-md-6">
                <label for="inputStatus">{{ label_status }}</label>
                <v-radio-group v-model="status">
                    <v-radio
                        label="Publish"
                        color="primary"
                        value="1"
                    ></v-radio>
                    <v-radio
                    label="Unpublish"
                    color="error"
                    value="0"
                    ></v-radio>
              </v-radio-group>
            </div>
        </div>
        <v-btn
        depressed
        color="primary"
        type="submit"
        >
            Submit
        </v-btn>
        <v-btn depressed @click="cancel(item)">
            Cancel
        </v-btn>
    </form>`
})

Vue.use(VeeValidate);
var app = new Vue({
    el: '#app',
    vuetify: new Vuetify(),
})