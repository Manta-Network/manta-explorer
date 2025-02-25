const path = require("path");

function resolve(dir) {
  return path.join(__dirname, dir);
}
export default {
  // mode: 'universal',
  mode: 'spa',
  /*
   ** Headers of the page
   */
  head: {
    title: process.env.npm_package_name || '',
    meta: [{
        charset: 'utf-8'
      },
      {
        name: 'viewport',
        content: 'width=device-width, initial-scale=1'
      },
      {
        hid: 'description',
        name: 'description',
        content: process.env.npm_package_description || ''
      }
    ],
    script: [{
      src: 'https://use.fontawesome.com/1f1cfc849c.js'
    }],
    link: [{
        rel: 'icon',
        type: 'image/x-icon',
        href: '/favicon.ico'
      },
      {
        rel: 'stylesheet',
        href: 'https://use.fontawesome.com/20cd1c1bd4.css'
      }
    ]
  },
  /*
   ** Customize the progress-bar color
   */
  loading: {
    color: '#fff'
  },
  /*
   ** Global CSS
   */
  css: [
    'element-ui/lib/theme-chalk/index.css'
  ],

  axios: {
    proxy: process.env.NODE_ENV !== 'production',
    browserBaseURL: process.env.NODE_ENV !== 'production' ? '' : 'https://api.e1.testnet.pelagosmanta.network'
  },

  // note: the proxy configuration below will not be used if running in production (on a host whose fqdn ends with 'manta.network').
  proxy: {
    "/api": {
      target: "http://127.0.0.1:4399",
      secure: false,
      changeOrigin: true,
      pathRewrite: {
        "^/api": "/api"
      }
    },
  },

  /*
   ** Plugins to load before mounting the App
   */
  plugins: [{
      src: '~/plugins/i18n.js'
    },
    {
      src: '~/plugins/icon'
    },
    {
      src: '~/plugins/element-ui'
    },
    {
      src: '~/plugins/axios'
    }
  ],
  /*
   ** Nuxt.js dev-modules
   */
  buildModules: [],
  /*
   ** Nuxt.js modules
   */
  modules: [
    '@nuxtjs/axios',
    '@nuxtjs/proxy'
  ],
  /*
   ** Build configuration
   */
  build: {
    transpile: [/^element-ui/, 'icon'],
    /*
     ** You can extend webpack config here
     */
    extend(config, ctx) {
      const svgRule = config.module.rules.find(rule => rule.test.test('.svg'))
      svgRule.exclude = [resolve('assets/icons')];

      // 添加svg-sprite-loader
      config.module
        .rules.push({
          test: (/\.svg$/),
          include: (resolve("assets/icons")),
          loader: ("svg-sprite-loader"),
          options: ({
            symbolId: "icon-[name]"
          })
        })
    }
  }
}
