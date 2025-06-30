
import { defineStore } from "pinia";
import { monitor } from "../../wailsjs/go/models";
 

export const useCtxConfigStore = defineStore('ctxConfig', {
    state: (): {
        selectedPath: string,
        selectedReq: string,



        c2s: Array<string>,
        c2sKeys: Array<string>,


        filaPaths: Array<string>,


        configData: Record<string, {
            keyPath: string,
            keyRules: Record<string, Array<monitor.Rule>>
        }>
    } => ({
        configData: {
            
        },
        selectedPath: '',
        selectedReq: '',
        filaPaths: ["e", "ew"],
        c2s: ["login", "do"], // 返回单条协议
        c2sKeys: ["a", "v"],
    }),

    actions: {
        load() {


        },

        AddConfig() {

            this.configData[this.selectedPath] = {}
        },
        AddRule() {
            if(! this.configData[this.selectedPath] ){
                this.configData[this.selectedPath]= {}
                this.configData[this.selectedPath].keyRules  = {}
                 this.configData[this.selectedPath].keyRules[this.selectedReq] = []
            }
            this.configData[this.selectedPath].keyRules[this.selectedReq].push({})
        },
        AddFilter(index: number) {
            if (!this.configData[this.selectedPath].keyRules[this.selectedReq][index].Fs) {
                this.configData[this.selectedPath].keyRules[this.selectedReq][index].Fs = []
            }
            this.configData[this.selectedPath].keyRules[this.selectedReq][index].Fs.push({})
        },
        Upload(){
            console.log(this.configData)
        }

    }
}) 