
import { defineStore } from 'pinia'
import { biz } from '../../wailsjs/go/models'
import { GetColumns, Pull, Start } from '../../wailsjs/go/biz/App'

// interface m {
//     columns: Array<biz.Column>,
//     c2s: Array<biz.Msg>
// }
let timerId;
export const useMonitorStore = defineStore('monitor', {
    state: (): {
        columns: Array<biz.Column>,
        c2s: Array<biz.Msg>
        s2c: Array<biz.Msg>
    } => ({
        columns: [],
        c2s: [],
        s2c: []

    }),
    actions: {
        GetColumns() {
            GetColumns().then((resp) => {
                console.log("resp ", resp)
                this.columns = resp
            })
        },
        Start() {
            Start(10002).then((resp) => {
                console.log(resp)

                timerId = setInterval(() => {


                    Pull().then((resp) => {
                        for (let k in resp) {
                            if (resp[k].is_client) {
                                this.c2s.push(resp[k])
                            } else {
                                this.s2c.push(resp[k])
                            }
                        }

                    })
                }, 5000)

            })
        },

        Stop() {
            clearInterval(timerId)
        },
        Pull() {
            Pull().then((resp) => {
                for (let k in resp) {
                    if (resp[k].is_client) {
                        this.c2s.push(resp[k])
                    } else {
                        this.s2c.push(resp[k])
                    }
                }
            })
        }
    }
})