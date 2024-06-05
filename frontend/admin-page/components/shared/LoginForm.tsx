import { loginFormSchema } from "@/lib/validator";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { Form, FormControl } from "../ui/form";
import { Input } from "../ui/input";

const LoginForm = () => {
  const initialValues = {
    email: "",
    password: "",
  };

  const form = useForm<z.infer<typeof loginFormSchema>>({
    defaultValues: initialValues,
    resolver: zodResolver(loginFormSchema),
  });

  const onSubmit = (data: z.infer<typeof loginFormSchema>) => {
    console.log(data);
  };

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-6">
        <FormControl>
          <Input
            id="email"
            type="email"
            placeholder="Email"
            {...form.register("email")}
          />
        </FormControl>
      </form>
    </Form>
  );
};
