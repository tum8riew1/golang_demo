const app_url = 'http://localhost:3000';
const app_path_dashboard = '/admin/dashboard';
const app_path_user_session = '/user-session';

const api_url = 'http://localhost:3001'
const api_path_login = '/api/v1/backend/login'


Vue.component('login-components', {
    mounted() {
        // axios
        //     .get(api_url + path_hashpassword)
        //     .then((response) => {
        //         console.log(response)
        //     })
    },
    props: {
        text_username: '',
        text_username_placeholder: '',
        text_password_placeholder: '',
        text_password: '',
        text_submit: '',
        validate_message_username_require: '',
        validate_message_password_require: '',
        url_user_session: app_url + app_path_user_session,
    },
    data() {
        return {
            username: '',
            password: '',
        }
    },
    methods: {
        api_login: async function(e) {
            console.log("Summit");
            //console.log(this.language);
            //console.log(this.parent_id.code);

            this.errors = [];

            if (!this.password) {
                this.errors.push(this.validate_message_password_require);
                this.$refs.password.focus();
            }


            if (!this.username) {
                this.errors.push(this.validate_message_username_require);
                this.$refs.username.focus();
            }

            if (!this.errors.length) {
                console.log("Case True");
                const params_login = {
                    username: this.username,
                    password: this.password
                };

                response_login = await axios.post(api_url + api_path_login, params_login, {
                    headers: {
                        'Content-Type': 'application/json',
                    },
                }).catch(function(error) {
                    console.log('Error! Could not reach the API. ' + error);
                });

                if (response_login.status == 200) {
                    console.log(response_login.status, response_login.data);
                    const params_user_session = {
                        token: response_login.data.token
                    };
                    response_user_session = await axios.post(app_url + app_path_user_session, params_user_session, {
                        headers: {
                            'Content-Type': 'application/json',
                        },
                    }).catch(function(error) {
                        console.log('Error! Could not reach the API. ' + error);
                    });
                    console.log(response_user_session.status);
                    if (response_user_session.status == 200) {
                        console.log(response_user_session.data);
                        location.href = app_url + app_path_dashboard;
                    } else {
                        console.log("session failed");
                    }
                }

                // axios.post(api_url + api_path_login, params, {
                //     headers: {
                //         'Content-Type': 'application/json',
                //     },
                // }).then(response => {

                //     console.log(response);
                //     //loader.hide();
                //     if (response.status == 200) {
                //         //loader.hide();
                //         // Vue.swal({
                //         //   position: 'center',
                //         //   type: 'success',
                //         //   title:this.message_save_data,
                //         //   showConfirmButton: false,
                //         //   timer: 5000
                //         // });
                //         console.log(app_url, app_path_dashboard)
                //             //location.href = app_url + app_path_dashboard;
                //     }

                // }).catch(function(error) {
                //     console.log('Error! Could not reach the API. ' + error);
                // });

                //console.log("Success");
                // e.preventDefault();
                // return false;
            }

            e.preventDefault();
        }
    },
    template: `<div class="container">
                    <div class="row align-items-center justify-content-center">
                        <div class="col-md-7">
                            <form @submit.prevent="api_login">
                                <h3>Login to <strong>Golang TuM8Riew</strong></h3>
                                <p class="mb-4"></p>
                                <div class="form-group first">
                                    <label for="username">{{ text_username }}</label>
                                    <input type="text" class="form-control" v-model="username" ref="username" v-bind:placeholder="text_username_placeholder" >
                                </div>
                                <div class="form-group last mb-3">
                                    <label for="password">{{ text_password }}</label>
                                    <input type="password" class="form-control"  v-model="password" ref="password" v-bind:placeholder="text_password_placeholder">
                                </div>
                                
                                <div class="d-flex mb-5 align-items-center">
                                    <label class="control control--checkbox mb-0"><span class="caption">Remember me</span>
                                    <input type="checkbox" checked="checked"/>
                                    <div class="control__indicator"></div>
                                    </label>
                                    <span class="ml-auto"><a href="#" class="forgot-pass">Forgot Password</a></span>
                                </div>

                                <button @click="api_login"  class="btn btn-block btn-primary">{{ text_submit }}</button>
                            </form>
                        </div>
                    </div>
                </div>`
})

var app = new Vue({
    el: '#app',
    vuetify: new Vuetify(),
})