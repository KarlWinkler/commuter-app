

export default {
  "meta": {},
  "id": "_default",
  "name": "",
  "file": {
    "path": "src/routes/_module.svelte",
    "dir": "src/routes",
    "base": "_module.svelte",
    "ext": ".svelte",
    "name": "_module"
  },
  "asyncModule": () => import('../src/routes/_module.svelte'),
  "rootName": "default",
  "routifyDir": import.meta.url,
  "children": [
    {
      "meta": {},
      "id": "_default_footer_svelte",
      "name": "footer",
      "file": {
        "path": "src/routes/footer.svelte",
        "dir": "src/routes",
        "base": "footer.svelte",
        "ext": ".svelte",
        "name": "footer"
      },
      "asyncModule": () => import('../src/routes/footer.svelte'),
      "children": []
    },
    {
      "meta": {},
      "id": "_default_header_svelte",
      "name": "header",
      "file": {
        "path": "src/routes/header.svelte",
        "dir": "src/routes",
        "base": "header.svelte",
        "ext": ".svelte",
        "name": "header"
      },
      "asyncModule": () => import('../src/routes/header.svelte'),
      "children": []
    },
    {
      "meta": {
        "isDefault": true
      },
      "id": "_default_index_svelte",
      "name": "index",
      "file": {
        "path": "src/routes/index.svelte",
        "dir": "src/routes",
        "base": "index.svelte",
        "ext": ".svelte",
        "name": "index"
      },
      "asyncModule": () => import('../src/routes/index.svelte'),
      "children": []
    },
    {
      "meta": {},
      "id": "_default_log_trip",
      "name": "log-trip",
      "module": false,
      "file": {
        "path": "src/routes/log-trip",
        "dir": "src/routes",
        "base": "log-trip",
        "ext": "",
        "name": "log-trip"
      },
      "children": [
        {
          "meta": {
            "isDefault": true
          },
          "id": "_default_log_trip_index_svelte",
          "name": "index",
          "file": {
            "path": "src/routes/log-trip/index.svelte",
            "dir": "src/routes/log-trip",
            "base": "index.svelte",
            "ext": ".svelte",
            "name": "index"
          },
          "asyncModule": () => import('../src/routes/log-trip/index.svelte'),
          "children": []
        }
      ]
    },
    {
      "meta": {},
      "id": "_default_success",
      "name": "success",
      "module": false,
      "file": {
        "path": "src/routes/success",
        "dir": "src/routes",
        "base": "success",
        "ext": "",
        "name": "success"
      },
      "children": [
        {
          "meta": {
            "isDefault": true
          },
          "id": "_default_success_index_svelte",
          "name": "index",
          "file": {
            "path": "src/routes/success/index.svelte",
            "dir": "src/routes/success",
            "base": "index.svelte",
            "ext": ".svelte",
            "name": "index"
          },
          "asyncModule": () => import('../src/routes/success/index.svelte'),
          "children": []
        }
      ]
    },
    {
      "meta": {
        "dynamic": true,
        "dynamicSpread": true,
        "order": false,
        "inline": false
      },
      "name": "[...404]",
      "file": {
        "path": ".routify/components/[...404].svelte",
        "dir": ".routify/components",
        "base": "[...404].svelte",
        "ext": ".svelte",
        "name": "[...404]"
      },
      "asyncModule": () => import('./components/[...404].svelte'),
      "children": []
    }
  ]
}