import { defineStore } from 'pinia'
import type { Playback, Req, Message } from '../types/playback'

export const usePlaybackStore = defineStore('playback', {
  state: (): Playback => ({
    c2s: [],
    s2c: [],
  }),
  actions: {
    addC2S(req: Req) {
      this.c2s.push(req)
    },
    addS2C(res: Message) {
      this.s2c.push(res)
    },
    clearAll() {
      this.c2s = []
      this.s2c = []
    },
     loadMockData() {
       this.c2s = [
        { time: 1749389602, name: 'LoginReq', data: '{"username":"u1"}' },
        { time: 1749389606, name: 'LoginReq', data: '{"username":"u2"}' },
        { time: 1749389612, name: 'MoveReq', data: '{"direction":"left"}' },
         { time: 1749389802, name: 'LoginReq', data: '{"username":"u1"}' },
        { time: 1749389906, name: 'LoginReq', data: '{"username":"u2"}' },
        { time: 1749389912, name: 'MoveReq', data: '{"direction":"left"}' },
      ]

      this.s2c = [
        { time: 1749389602, name: 'LoginResp', data: '{"data":{"user":{"id":101},"token":"abc"}}' },
        { time: 1749389602, name: 'MoveResp', data: '{"pos":{"x":100,"y":200}}' },
         { time: 1749389802, name: 'LoginResp', data: '{"data":{"user":{"id":101},"token":"abc"}}' },
        { time: 1749389802, name: 'MoveResp', data: '{"pos":{"x":100,"y":200}}' },
         { time: 1749389412, name: 'LoginResp', data: '{"data":{"user":{"id":101},"token":"abc"}}' },
        { time: 1749389312, name: 'MoveResp', data: '{"pos":{"x":100,"y":200}}' },
         { time: 1749389915, name: 'LoginResp', data: '{"data":{"user":{"id":101},"token":"abc"}}' },
        { time: 17493899123, name: 'MoveResp', data: '{"pos":{"x":100,"y":200}}' },
      ]



     }
  }
})
