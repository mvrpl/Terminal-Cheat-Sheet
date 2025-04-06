<script lang="ts" setup>
import {reactive} from 'vue'
import {Greet} from '../../wailsjs/go/main/App'

const data = reactive({
  name: "",
  resultText: "Please enter your name below 👇",
})

function greet() {
  Greet(data.name).then(result => {
    data.resultText = result
  })
}

</script>

<template>
  <div class="content-area">
    <h2 id="tabela1">Tabela 1</h2>
        <table>
            <tr>
                <td>Dados 1</td>
                <td>Dados 2</td>
            </tr>
        </table>

        <h2 id="tabela2">Tabela 2</h2>
        <table>
            <tr>
                <td>Dados A</td>
                <td>Dados B</td>
            </tr>
        </table>

        <h2 id="tabela3">Tabela 3</h2>
        <table>
            <tr>
                <td>Dados X</td>
                <td>Dados Y</td>
            </tr>
        </table>

        <h2 id="tabela3">Tabela 3</h2>
        <table>
            <tr>
                <td>Dados X</td>
                <td>Dados Y</td>
            </tr>
        </table>

        <h2 id="tabela3">Tabela 3</h2>
        <table>
            <tr>
                <td>Dados X</td>
                <td>Dados Y</td>
            </tr>
        </table>

        <h2 id="tabela3">Tabela 3</h2>
        <table>
            <tr>
                <td>Dados X</td>
                <td>Dados Y</td>
            </tr>
        </table>

        <h2 id="tabela3">Tabela 3</h2>
        <table>
            <tr>
                <td>Dados X</td>
                <td>Dados Y</td>
            </tr>
        </table>
  </div>
</template>

<style scoped>
.content-area {
  margin-left: 220px;
  padding: 20px;
  width: calc(100% - 220px);
}
</style>