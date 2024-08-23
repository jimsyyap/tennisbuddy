<template>
    <header>
        <h1>
            <span class="icon" aria-hidden="true">
                check &#x2705;
            </span>
        </h1>
        <p>
            Lorem ipsum dolor sit amet consectetur adipisicing elit. A necessitatibus, maxime iste rerum fugit fugiat incidunt nisi reiciendis, eaque quisquam molestiae totam porro voluptates! Nihil recusandae esse aliquid placeat quis!
        </p>
    </header>
</template>

<style scoped>
header {
    border-bottom: #333 1px solid;
    background-color: #fff;
}
header::after {
    content: '';
    display: block;
    height: 1px;
    box-shadown: 0px 0px 10px 0px rgba(0, 0, 0, 0.75);
}
h1 {
    font-size: 2em;
}
h1 .icon {
    font-size: 1rem;
    vertical-align: middle;
}
</style>

<script>
</script>
