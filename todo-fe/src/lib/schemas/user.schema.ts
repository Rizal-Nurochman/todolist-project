import { z } from "zod";
export class UserSchema{
    static create=z.object({
        email:z.string().email(),
        password:z.string().min(8),
        confirmPassword:z.string()
    }).refine((data) => data.password === data.confirmPassword, {
        message: "Passwords don't match",
        path: ["confirmPassword"],
    });
}

export type CreateUserSchema = z.infer<typeof UserSchema.create>