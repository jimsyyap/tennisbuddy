<script lang="ts" setup> 
    import ListItem from './ListItem.vue' 
    type Item = {
        title: string,
        checked?: boolean
    }
    const items: listItems: Item[] = [
        { title: 'Buy milk', checked: false },
        { title: 'Buy eggs', checked: true }
    ]
</script>

<template>
    <ul>
        <li
            :key='key'
            v-for='(item, key) in items'
        >
        <ListItem :is-checked='item.checked'>
            {{ item.title }}
        </ListItem>

    </ul>
</template>
