func asteroidCollision(asteroids []int) []int {
    stack := []int{}

    for _, ast := range asteroids {
        destroyed := false

        // While there is a possible collision (right-moving asteroid on stack, left-moving current asteroid)
        for len(stack) > 0 && stack[len(stack)-1] > 0 && ast < 0 {
            top := stack[len(stack)-1]

            if top < -ast {
                // Top asteroid is smaller -> pop it and check next one
                stack = stack[:len(stack)-1]
                continue
            } else if top == -ast {
                // Both are equal -> pop top, destroy current
                stack = stack[:len(stack)-1]
            }
            // In both equal or larger cases, current asteroid is destroyed
            destroyed = true
            break
        }

        if !destroyed {
            stack = append(stack, ast)
        }
    }

    return stack
}
