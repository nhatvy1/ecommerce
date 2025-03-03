import { Controller, Get, Param } from '@nestjs/common'
import { CategoryService } from './category.service';

@Controller('category')
export class CategoryController {
  constructor(private readonly categoryService: CategoryService) {}

  @Get('id')
  findById(@Param('id') id: string) {
    return this.categoryService.findById(id)
  }
}
