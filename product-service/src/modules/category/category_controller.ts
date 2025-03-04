import { Body, Controller, Delete, Get, HttpCode, HttpStatus, Param, Patch, Post } from '@nestjs/common'
import { CategoryService } from './category_service'
import { ResponseMessage } from 'src/shared/decorators/response_message_decorator'
import { CreateCategoryDto } from './dto/create_category_dto'
import { UpdateCategoryDto } from './dto/update_category_dto'
import { IdDto } from './dto/id_dto'

@Controller('category')
export class CategoryController {
  constructor(private readonly categoryService: CategoryService) {}

  @Get('id')
  @ResponseMessage("success")
  @HttpCode(HttpStatus.OK)
  findById(@Param('id') idDto: IdDto) {
    return this.categoryService.findById(idDto.id)
  }
  
  @Post('')
  @ResponseMessage("success")
  @HttpCode(HttpStatus.CREATED)
  createCategory(@Body() body: CreateCategoryDto) {
    return this.categoryService.create(body)
  }

  @Patch('id')
  @ResponseMessage("success")
  @HttpCode(HttpStatus.OK)
  updateCategory(@Param('id') idDto: IdDto, @Body() body: UpdateCategoryDto) {
    return this.categoryService.update(idDto.id, body)
  }

  @Delete('id')
  @ResponseMessage("success")
  @HttpCode(HttpStatus.OK)
  deleteCategory(@Param('id') idDto: IdDto) {
    return this.categoryService.delete(idDto.id)
  }
}
