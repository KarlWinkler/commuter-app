<script lang="ts" context="module">
  import { Router, createRouter } from '@roxi/routify'
  import routes from '../.routify/routes.default.js'

  import './styles/app.css'

  export const router = createRouter({ routes })
</script>

<Router {router} />
